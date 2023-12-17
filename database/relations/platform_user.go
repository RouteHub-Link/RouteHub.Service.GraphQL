package database_relations

type PlatformUser struct {
	UserID     string  `json:"userId"`
	PlatformID string  `json:"platformId"`
	Role       string  `json:"role"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
	DeletedAt  *string `json:"deletedAt,omitempty"`
}
