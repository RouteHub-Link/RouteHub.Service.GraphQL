package database_models

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type Link struct {
	ID                uuid.UUID                         `gorm:"type:uuid;primary_key;"`
	PlatformID        uuid.UUID                         `gorm:"type:uuid;not null;field:platform_id"`
	Target            string                            `gorm:"type:varchar(255);not null;"`
	Path              string                            `gorm:"type:varchar(255);not null;"`
	OpenGraph         *database_types.OpenGraph         `json:"openGraph" gorm:"serializer:json"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice" gorm:"serializer:json"`
	Status            database_enums.StatusState        `json:"status" gorm:"serializer:json"`

	CreatedBy uuid.UUID  `gorm:"type:uuid;not null;"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `gorm:"index"`
}
