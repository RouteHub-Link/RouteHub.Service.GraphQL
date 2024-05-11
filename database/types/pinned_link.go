package database_types

import "github.com/google/uuid"

type PinnedLink struct {
	LinkID    uuid.UUID `json:"linkId"`
	PinnedBy  uuid.UUID `json:"pinnedBy"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

func (PinnedLink) TableName() string {
	return "pinned_links"
}
