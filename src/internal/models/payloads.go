package models

type QueryPayload struct {
	Skip *int64 `json:"skip,omitempty"`
	Take *int64 `json:"take,omitempty"`
}
