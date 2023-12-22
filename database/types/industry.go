package database_types

import "time"

type Industry struct {
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
