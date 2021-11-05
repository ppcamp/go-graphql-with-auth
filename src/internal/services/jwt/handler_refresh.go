package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-graphql-with-auth/internal/config"
	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
)

// RefreshHandler get the current session and generates a new jwt token.
// Note that this function doesn't care about unvalidate the old one.
func (md *JWTMiddleware) RefreshHandler(c *gin.Context) (loginmodels.TokenResponse, error) {
	session, err := GetSession(c)
	if err != nil {
		return loginmodels.TokenResponse{}, errors.New("fail to get the current session")
	}

	return buildAndResponseToken(
		c,
		session,
		config.App.ApiDomain,
		config.App.JWTExp,
		[]byte(config.App.JWTSecret))
}
