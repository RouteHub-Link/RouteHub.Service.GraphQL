package database_relations

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationPlatform struct {
	ID             uuid.UUID  `gorm:"primaryKey;type:uuid;not null"`
	OrganizationID uuid.UUID  `json:"organizationId"`
	PlatformID     uuid.UUID  `json:"platformId"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
}
