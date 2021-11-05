package usermodels

import "time"

type UserEntity struct {
	Id        *int64     `json:"id,omitempty" db:"id"`
	Password  *string    `json:"password,omitempty" db:"password"`
	Nick      *string    `json:"nick,omitempty" db:"nick"`
	Email     *string    `json:"email,omitempty" db:"email"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
