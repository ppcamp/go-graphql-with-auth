package auth

import (
	"errors"

	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
)

var (
	ErrInvalidPayload      = errors.New("invalid payload")
	ErrKeyValue            = errors.New("invalid payload infos")
	ErrInvalidLoginPayload = errors.New("invalid login payload")
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// ApiAuthenticator does the login checkup.
// If everything occurrs well, it'll returns the JWT session to be encoded into
// JWT
func (a *AuthService) ApiAuthenticator(login *loginmodels.LoginPayload) (jwt.Session, error) {
	if login.Password == "coxinha" {
		return jwt.Session{
			LoginSession: jwt.LoginSession{},
		}, nil
	}

	return jwt.BLANK_SESSION, ErrInvalidPayload
}
