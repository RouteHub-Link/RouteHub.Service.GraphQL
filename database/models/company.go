package database_models

import (
	"time"

	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type Company struct {
	ID           uuid.UUID                     `json:"id" gorm:"primaryKey;type:uuid;field:id"`
	Name         string                        `json:"name"`
	Website      string                        `json:"website"`
	Industry     []*database_types.Industry    `json:"industry,omitempty" gorm:"serializer:json"`
	Description  string                        `json:"description"`
	Location     string                        `json:"location"`
	SocialMedias []*database_types.SocialMedia `json:"socialMedias" gorm:"serializer:json"`

	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
