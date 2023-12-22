package database_models

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	OrganizationId uuid.UUID `gorm:"type:uuid;not null;"`
	Name           string    `gorm:"type:varchar(255);not null;"`
	URL            string    `gorm:"type:varchar(255);not null;"`

	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

/*
type Domain struct {
	ID                    uuid.UUID                     `json:"id"`
	Name                  string                        `json:"name"`
	URL                   string                        `json:"url"`
	Organization          *database_models.Organization `json:"organization"`
	Platform              *database_models.Platform     `json:"platform,omitempty"`
	Verification          []*DomainVerification         `json:"verification"`
	State                 StatusState                   `json:"state"`
	Links                 []*Link                       `json:"links,omitempty"`
	Analytics             []*MetricAnalytics            `json:"analytics"`
	AnalyticReports       *AnalyticReports              `json:"analyticReports"`
	LastDNSVerificationAt *time.Time                    `json:"lastDNSVerificationAt,omitempty"`
	CreatedAt             time.Time                     `json:"createdAt"`
	UpdatedAt             *time.Time                    `json:"updatedAt,omitempty"`
	DeletedAt             *time.Time                    `json:"deletedAt,omitempty"`
}
*/
