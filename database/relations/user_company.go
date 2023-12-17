package database_relations

import "github.com/google/uuid"

type UserCompany struct {
	UserID    uuid.UUID `json:"user_id" gorm:"field:user_id;primaryKey;type:uuid;not null"`
	CompanyID uuid.UUID `json:"company_id"  gorm:"field:company_id;primaryKey;type:uuid;not null"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt *string   `json:"updatedAt,omitempty"`
	DeletedAt *string   `json:"deletedAt,omitempty"`
}
