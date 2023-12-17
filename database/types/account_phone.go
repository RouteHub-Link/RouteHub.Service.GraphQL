package database_types

type AccountPhone struct {
	Number      string  `json:"number"`
	CountryCode string  `json:"countryCode"`
	Verified    bool    `json:"verified"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
}
