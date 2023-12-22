package database_types

import "time"

type AccountPhone struct {
	Number      string     `json:"number"`
	CountryCode string     `json:"countryCode"`
	Verified    bool       `json:"verified"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
