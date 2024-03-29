package database_relations

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationPlatform struct {
	ID             uuid.UUID  `gorm:"primaryKey;type:uuid;not null"`
	OrganizationID uuid.UUID  `json:"organization_id" gorm:"field:organization_id;type:uuid;not null"`
	PlatformID     uuid.UUID  `json:"platform_id" gorm:"field:platform_id;type:uuid;not null"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
}
