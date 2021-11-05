package jwt

import (
	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
)

type ApiAuthenticator func(login *loginmodels.LoginPayload) (Session, error)
