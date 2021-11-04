package controller

type JsonError struct {
	Status int         `json:"status,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}
