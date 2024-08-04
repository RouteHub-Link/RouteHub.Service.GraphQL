package database_models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;field:id"`
	Subject   string     `json:"subject" gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
