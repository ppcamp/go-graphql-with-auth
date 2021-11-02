package models

import "time"

type User struct {
	Id        *int64     `json:"id,omitempty" db:"id"`
	Password  *string    `json:"password,omitempty" db:"password" binding:"min=3"`
	Nick      *string    `json:"nick,omitempty" db:"nick"`
	Email     *string    `json:"email,omitempty" db:"email"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
