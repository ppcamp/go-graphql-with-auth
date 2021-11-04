package usermodels

import (
	"time"

	"github.com/ppcamp/go-graphql-with-auth/internal/models"
)

type UserMutationPayload struct {
	Id        *int64     `json:"id,omitempty" db:"id" binding:"omitempty"`
	Password  *string    `json:"password,omitempty" db:"password" binding:"omitempty,min=3"`
	Nick      *string    `json:"nick,omitempty" db:"nick" binding:"omitempty"`
	Email     *string    `json:"email,omitempty" db:"email" binding:"omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at" binding:"omitempty"`
}

type UserQueryPayload struct {
	Id        *int64     `json:"id,omitempty" db:"id" binding:"omitempty"`
	Nick      *string    `json:"nick,omitempty" db:"nick" binding:"omitempty"`
	Email     *string    `json:"email,omitempty" db:"email" binding:"omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at" binding:"omitempty"`

	*models.QueryPayload
}
