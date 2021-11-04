package models

type QueryPayload struct {
	Skip *int64 `json:"skip,omitempty" db:"skip,omitempty"`
	Take *int64 `json:"take,omitempty" db:"take,omitempty"`
}
