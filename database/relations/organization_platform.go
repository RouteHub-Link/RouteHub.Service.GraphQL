package database_relations

type OrganizationPlatform struct {
	OrganizationID string  `json:"organizationId"`
	PlatformID     string  `json:"platformId"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      *string `json:"updatedAt,omitempty"`
	DeletedAt      *string `json:"deletedAt,omitempty"`
}
