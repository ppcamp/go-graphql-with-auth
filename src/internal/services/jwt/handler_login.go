package jwt

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/ppcamp/go-graphql-with-auth/internal/config"
// 	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
// )

// // LoginHandler will validate the login and, if the login is a valid one,
// // returns a jwt signed token.
// // See auth.ApiAuthenticator
// func (md *JWTMiddleware) LoginHandler(c *gin.Context) {
// 	login := &loginmodels.LoginPayload{}

// 	err := c.ShouldBindJSON(&login)
// 	if err != nil {
// 		// c.Error(ErrInvalidLoginPayload)
// 		// c.AbortWithStatus(http.StatusBadRequest)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.JsonError{
// 			Status: http.StatusBadRequest,
// 			Errors: validators.MapErrorsToMessages(err),
// 		})
// 		return
// 	}

// 	session, err := md.ApiAuthenticator(login)
// 	if err != nil {
// 		// c.Error(err)
// 		// c.AbortWithStatus(http.StatusUnauthorized)
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.JsonError{
// 			Status: http.StatusUnauthorized,
// 			Errors: []string{ErrInvalidLoginPayload.Error()},
// 		})
// 		return
// 	}

// 	md.buildAndResponseToken(c, session, config.App.ApiDomain)
// }
