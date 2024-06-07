package database_models

import (
	"time"

	"github.com/google/uuid"
)

type LinkValidation struct {
	ID            uuid.UUID  `json:"id"`
	LinkId        uuid.UUID  `json:"link" gorm:"type:uuid;not null;field:link_id"`
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

func (LinkValidation) TableName() string {
	return "link_validations"
}

func (lv *LinkValidation) Requested(link *Link, userId uuid.UUID, taskId string) {
	lv.ID = uuid.New()
	lv.TaskId = taskId
	lv.LinkId = link.ID
	lv.CreatedBy = userId
}

func (lv *LinkValidation) Ended(isValid bool, message *string, err *string, completedAt time.Time) {
	lv.IsValid = isValid
	lv.Message = message
	lv.Error = err
	lv.CompletedAt = &completedAt

	updatedAt := time.Now()
	lv.UpdatedAt = &updatedAt
}
