package database_models

import (
	"time"

	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID                    `json:"id" gorm:"primaryKey;type:uuid;field:id"`
	Avatar       string                       `json:"avatar"`
	Email        string                       `json:"email"`
	Fullname     string                       `json:"fullname"`
	Verified     bool                         `json:"verified"`
	Phone        *database_types.AccountPhone `json:"phone" gorm:"serializer:json"`
	PasswordHash string                       `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt    time.Time                    `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time                   `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt    *time.Time                   `json:"deletedAt,omitempty" gorm:"index"`
}
