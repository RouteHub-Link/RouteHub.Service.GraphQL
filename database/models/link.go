package database_models

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

//  link.TargetDetails String by default html for redirection custom

type Link struct {
	ID                uuid.UUID                         `gorm:"type:uuid;primary_key;"`
	PlatformID        uuid.UUID                         `gorm:"type:uuid;not null;field:platform_id"`
	Platform          *Platform                         `json:"platform" gorm:"foreignKey:PlatformID;references:id"`
	Target            string                            `gorm:"type:varchar(255);not null;"`
	Path              string                            `gorm:"type:varchar(255);not null;"`
	LinkContent       *database_types.LinkContent       `json:"link_content" gorm:"serializer:json;field:link_content"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice" gorm:"serializer:json"`
	State             database_enums.StatusState        `json:"State" gorm:"serializer:json"`
	ValidationTaskID  uuid.UUID                         `gorm:"type:uuid;field:validation_task_id"`

	CreatedBy uuid.UUID  `gorm:"type:uuid;not null;"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `gorm:"index"`
}

func (Link) TableName() string {
	return "links"
}
