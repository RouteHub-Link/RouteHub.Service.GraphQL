package database_relations

type CompanyPlatform struct {
	CompanyID  string  `json:"companyId"`
	PlatformID string  `json:"platformId"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
	DeletedAt  *string `json:"deletedAt,omitempty"`
}
