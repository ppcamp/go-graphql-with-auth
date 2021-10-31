package controller

import (
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
)

type Handler struct {
	Storage postgres.Storage
}

func NewControllersHandler() *Handler {
	return &Handler{}
}
