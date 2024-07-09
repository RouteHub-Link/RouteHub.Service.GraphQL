package services_domain

import (
	"encoding/base64"
	"errors"
	"log"
	"time"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	services_domain_utils "github.com/RouteHub-Link/routehub-service-graphql/services/domain_utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DNSVerificationService struct {
	domainUtilsService *services_domain_utils.DomainUtilsService
	db                 *gorm.DB
}

func NewDNSVerificationService(domainUtilsService *services_domain_utils.DomainUtilsService, db *gorm.DB) *DNSVerificationService {
	return &DNSVerificationService{domainUtilsService: domainUtilsService, db: db}
}
func (dvs *DNSVerificationService) IsVerificationProcessing(domain *database_models.Domain) bool {
	dnsVerification := &database_models.DNSVerification{}
	err := dvs.db.Where("domain_id = ? AND completed_at IS NULL", domain.ID).Last(&dnsVerification).Error
	if err != nil {

		return false
	}

	return true
}

func (dvs *DNSVerificationService) Validate(userId uuid.UUID, domain *database_models.Domain, force bool) (dnsVerification *database_models.DNSVerification, err error) {
	activeDNSVerification := &database_models.DNSVerification{}

	err = dvs.db.Where("domain_id = ? AND completed_at IS NULL", domain.ID).Last(&activeDNSVerification).Error

	if err == nil {
		if !force {
			return nil, errors.New("there is an active DNS verification")
		}

		cancelledMessage := "Forced to cancel"

		activeDNSVerification.Cancelled(&cancelledMessage, nil)
		err = dvs.db.Save(activeDNSVerification).Error
		if err != nil {
			return
		}
	}

	generatedSecret := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	dnsVerification = &database_models.DNSVerification{}

	payload := &services_domain_utils.DNSVerificationPayload{Link: domain.URL, Value: generatedSecret}
	taskId, err := dvs.domainUtilsService.PostValidateDNS(payload)
	if err != nil {
		return
	}

	dnsVerification.Requested(domain, userId, taskId, generatedSecret)

	err = dvs.db.Create(dnsVerification).Error
	return
}

func (dvs *DNSVerificationService) GetDNSVerificationByDomain(domain *database_models.Domain) (dnsVerification *database_models.DNSVerification, err error) {
	return dvs.GetDNSVerificationByDomainId(domain.ID)
}

func (dvs *DNSVerificationService) GetDNSVerificationByDomainId(domainId uuid.UUID) (dnsVerification *database_models.DNSVerification, err error) {
	dnsVerification = &database_models.DNSVerification{}
	err = dvs.db.Where("domain_id = ?", domainId).Last(&dnsVerification).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return
	}

	if dnsVerification == nil {
		log.Default().Printf("DNSVerification is nil %+v", domainId)
		return
	}

	err = checkIsCompleted(dnsVerification, dvs)
	return
}

func (dvs *DNSVerificationService) GetDNSVerificationsByDomainId(domainId uuid.UUID) (dnsVerifications []*database_models.DNSVerification, err error) {
	err = dvs.db.Where("domain_id = ?", domainId).Find(&dnsVerifications).Error
	if err != nil {
		return
	}

	for _, dnsVerification := range dnsVerifications {
		err = checkIsCompleted(dnsVerification, dvs)
		if err != nil {
			log.Default().Printf("\nError checking DNSVerification %+v\n dnsVerificationRecord %+v", err, dnsVerification)
			return
		}
	}

	return
}

func (dvs *DNSVerificationService) GetDNSVerificationsByDomain(domain *database_models.Domain) (dnsVerifications []*database_models.DNSVerification, err error) {
	return dvs.GetDNSVerificationsByDomainId(domain.ID)
}

func checkIsCompleted(dnsVerification *database_models.DNSVerification, dvs *DNSVerificationService) (err error) {
	if dnsVerification.CompletedAt != nil {
		return
	}

	if dnsVerification.TaskId == "" {
		return errors.New("TaskId is empty")
	}

	checkTime := dnsVerification.CreatedAt.Add(1 * time.Minute)
	now := time.Now()
	if dnsVerification.LastCheckedAt != nil && checkTime.After(now) {
		log.Default().Printf("Last checked time is less than 1 minutes %+v", dnsVerification)
		return nil
	}

	dnsVerification.LastCheckedAt = &now
	err = dvs.db.Save(dnsVerification).Error
	if err != nil {
		return
	}

	state, res, taskInfo, err := dvs.domainUtilsService.GetValidateDNS(dnsVerification.TaskId)
	if err != nil {
		return
	}

	if state == services_domain_utils.TaskStateCompleted {
		dnsVerification.Ended(res.IsValid, res.Message, res.Error, taskInfo.CompletedAt)
	} else {
		dnsVerification.NextProcessAt = &taskInfo.NextProcessAt
	}

	err = dvs.db.Save(dnsVerification).Error
	return
}
