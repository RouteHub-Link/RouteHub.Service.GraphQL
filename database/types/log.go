package database_types

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Message   string    `json:"message"`
}
