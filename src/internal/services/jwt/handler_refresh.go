package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-graphql-with-auth/internal/config"
)

// RefreshHandler get the current session and generates a new jwt token.
// Note that this function doesn't care about unvalidate the old one.
func (md *JWTMiddleware) RefreshHandler(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	md.buildAndResponseToken(c, session, config.App.ApiDomain)
}
