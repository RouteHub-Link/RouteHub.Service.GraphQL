package database_models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type DNSVerification struct {
	ID            uuid.UUID  `json:"id"`
	DomainId      uuid.UUID  `json:"domain" gorm:"type:uuid;not null;field:domain_id"`
	Secret        string     `json:"secret,omitempty" gorm:"not null;"`
	TaskId        string     `json:"target,omitempty" gorm:"not null;"`
	IsValid       bool       `json:"isValid,omitempty" gorm:"not null;"`
	Message       *string    `json:"message,omitempty" gorm:"serializer:json;"`
	Error         *string    `json:"error,omitempty" gorm:"serializer:json;"`
	CreatedBy     uuid.UUID  `gorm:"type:uuid;not null;"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	LastCheckedAt *time.Time `json:"lastChecked,omitempty" gorm:"field:last_checked_at;"`
	NextProcessAt *time.Time `json:"nextProcessAt,omitempty" gorm:"field:next_process_at;"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime:milli"`
	CompletedAt   *time.Time `gorm:"field:completed_at;"`
}

func (DNSVerification) TableName() string {
	return "dns_verifications"
}

func (dv *DNSVerification) Requested(domain *Domain, userId uuid.UUID, taskId string, secret string) {
	dv.ID = uuid.New()

	_taskId := dv.trimTaskIdString(taskId)

	dv.TaskId = _taskId
	dv.DomainId = domain.ID
	dv.CreatedBy = userId
	dv.Secret = secret
}

func (dv *DNSVerification) Ended(isValid bool, message *string, err *string, completedAt time.Time) {
	dv.IsValid = isValid
	dv.Message = message
	dv.Error = err
	dv.CompletedAt = &completedAt

	updatedAt := time.Now()
	dv.UpdatedAt = &updatedAt
}

func (dv *DNSVerification) Cancelled(message *string, err *string) {
	dv.IsValid = false
	dv.Message = message
	dv.Error = err

	updatedAt := time.Now()
	dv.UpdatedAt = &updatedAt
	dv.CompletedAt = &updatedAt
}

func (*DNSVerification) trimTaskIdString(taskId string) string {
	_taskId := taskId
	_taskId = strings.Trim(_taskId, "\n")
	_taskId = strings.Trim(_taskId, " ")
	_taskId = strings.Trim(_taskId, "\"")
	return _taskId
}
