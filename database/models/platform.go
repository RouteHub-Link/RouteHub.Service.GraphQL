package database_models

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type Platform struct {
	ID                uuid.UUID                         `gorm:"type:uuid;primary_key;"`
	Name              string                            `gorm:"type:varchar(255);not null;"`
	Slug              string                            `gorm:"type:varchar(255);not null;"`
	CreatedBy         uuid.UUID                         `gorm:"type:uuid;not null;"`
	DomainId          uuid.UUID                         `gorm:"type:uuid;not null;field:domain_id"`
	OpenGraph         *database_types.OpenGraph         `json:"openGraph" gorm:"serializer:json"`
	RedirectionChoice database_enums.RedirectionOptions `json:"redirectionChoice" gorm:"serializer:json"`
	Status            database_enums.StatusState        `json:"status" gorm:"serializer:json"`

	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

func (Platform) TableName() string {
	return "platforms"
}

/*
type Platform struct {
	ID                uuid.UUID                           `json:"id"`
	Name              string                              `json:"name"`
	OpenGraph         *OpenGraph                          `json:"openGraph"`
	RedirectionChoice database_enums.RedirectionOptions   `json:"redirectionChoice"`
	Organization      *database_models.Organization       `json:"organization"`
	Domains           []*Domain                           `json:"domains"`
	Permissions       []database_enums.PlatformPermission `json:"permissions"`
	Links             []*Link                             `json:"links"`
	Analytics         *PlatformAnalytics                  `json:"analytics"`
	Status            StatusState                         `json:"status"`
	Templates         []*Template                         `json:"templates"`
	PinnedLinks       []*Link                             `json:"pinnedLinks"`
}
*/
