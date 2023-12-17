package database_types

type Industry struct {
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
}
