package services_link

import (
	"errors"
	"log"
	"time"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	services_domain_utils "github.com/RouteHub-Link/routehub-service-graphql/services/domain_utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LinkValidationService struct {
	domainUtilsService *services_domain_utils.DomainUtilsService
	db                 *gorm.DB
}

func NewLinkValidationService(domainUtilsService *services_domain_utils.DomainUtilsService, db *gorm.DB) *LinkValidationService {
	return &LinkValidationService{domainUtilsService: domainUtilsService, db: db}
}

func (lvs *LinkValidationService) IsValidationProcessing(link *database_models.Link) bool {
	linkValidation := &database_models.LinkValidation{}
	err := lvs.db.Where("link_id = ? AND completed_at IS NULL", link.ID).Last(&linkValidation).Error
	if err != nil {
		return false
	}

	return true
}

func (lvs *LinkValidationService) Validate(userId uuid.UUID, link *database_models.Link) (linkValidation *database_models.LinkValidation, err error) {
	if lvs.IsValidationProcessing(link) {
		return nil, errors.New("Validation is already processing")
	}

	linkValidation = &database_models.LinkValidation{}
	payload := &services_domain_utils.SiteValidationPayload{Link: link.Target}
	taskId, err := lvs.domainUtilsService.PostValidateSite(payload)
	if err != nil {
		return
	}

	linkValidation.Requested(link, userId, taskId)
	log.Printf("LinkValidation.Requested %+v", linkValidation)

	err = lvs.db.Create(linkValidation).Error
	return
}

// GetLinkValidationByLink returns the last link validation for a given link if is not exist returns gorm.ErrRecordNotFound
func (lvs *LinkValidationService) GetLinkValidationByLink(link *database_models.Link) (linkValidation *database_models.LinkValidation, err error) {
	linkValidation = &database_models.LinkValidation{}
	err = lvs.db.Where("link_id = ?", link.ID).Last(&linkValidation).Error
	if err != nil {
		return
	}

	err = checkLinkValidation(linkValidation, lvs)
	return
}

// GetLinkValidationByLinkId returns the last link validation for a given link id if is not exist returns nil
func (lvs *LinkValidationService) GetLinkValidationByLinkId(linkId uuid.UUID) (linkValidation *database_models.LinkValidation, err error) {
	linkValidation = &database_models.LinkValidation{}
	err = lvs.db.Where("link_id = ?", linkId).Last(&linkValidation).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return
	}

	err = checkLinkValidation(linkValidation, lvs)
	return
}

// GetLinkValidationsByLinkId returns all link validations for a given link id sends the get validation request for every uncompleted validation requests. Its heavy operation.
func (lvs *LinkValidationService) GetLinkValidationsByLinkId(id uuid.UUID) (linkValidations []*database_models.LinkValidation, err error) {
	linkValidations = []*database_models.LinkValidation{}
	err = lvs.db.Where("link_id = ?", id).Find(&linkValidations).Error
	if err != nil {
		return
	}

	for _, linkValidation := range linkValidations {
		err = checkLinkValidation(linkValidation, lvs)
		if err != nil {
			log.Default().Printf("\nError checking link validation: %+v\n link validation: %+v ", err, linkValidation)
			return nil, err
		}
	}
	return
}

func checkLinkValidation(linkValidation *database_models.LinkValidation, lvs *LinkValidationService) (err error) {
	if linkValidation.CompletedAt != nil {
		return
	}

	if linkValidation.TaskId == "" {
		return errors.New("TaskId is empty")
	}

	checkTime := linkValidation.CreatedAt.Add(1 * time.Minute)
	now := time.Now()
	if linkValidation.LastCheckedAt != nil && checkTime.After(now) {
		log.Default().Printf("Last checked time is less than 1 minutes %+v", linkValidation)
		return nil
	}

	linkValidation.LastCheckedAt = &now
	err = lvs.db.Save(linkValidation).Error
	if err != nil {
		return
	}

	state, res, taskInfo, err := lvs.domainUtilsService.GetValidateSite(linkValidation.TaskId)
	if err != nil {
		return
	}

	if state == services_domain_utils.TaskStateCompleted {
		linkValidation.Ended(res.IsValid, res.Message, res.Error, taskInfo.CompletedAt)
	} else {
		linkValidation.NextProcessAt = &taskInfo.NextProcessAt
	}

	err = lvs.db.Save(linkValidation).Error
	return
}
