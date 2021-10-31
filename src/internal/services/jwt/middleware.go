package jwt

// import (
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/eventials/vlab-baby-app-api/internal/config"
// 	"github.com/eventials/vlab-baby-app-api/internal/helpers"
// 	"github.com/eventials/vlab-baby-app-api/internal/helpers/validators"
// 	loginmodels "github.com/eventials/vlab-baby-app-api/internal/models/login"
// 	"github.com/gin-gonic/gin"
// 	jwt "github.com/golang-jwt/jwt"
// )

// type JWTMiddleware struct {
// 	Expires          time.Duration
// 	Key              []byte
// 	ApiAuthenticator ApiAuthenticator
// }

// //#region: Generating JWT token

// // Token login response
// type tokenResponse struct {
// 	Token   string `json:"token,omitempty"`   // Session token
// 	Expires string `json:"expires,omitempty"` // Expiration timestamp
// }

// // buildAndResponseToken generates a signed endpoint with expiration
// func (md *JWTMiddleware) buildAndResponseToken(c *gin.Context, session Session, domain string) {
// 	exp := time.Now().UTC().Add(md.Expires)
// 	signedString, err := generateJwtAsSignedString(session, exp, md.Key)

// 	if err != nil {
// 		// c.Error(err)
// 		// c.AbortWithStatus(http.StatusUnauthorized)
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.JsonError{
// 			Status: http.StatusUnauthorized,
// 			Errors: []string{err.Error()},
// 		})
// 		return
// 	}

// 	c.SetCookie("session", signedString, int(md.Expires.Seconds()), "/", domain, false, false)

// 	c.JSON(http.StatusOK, tokenResponse{
// 		Token:   signedString,
// 		Expires: exp.Format(time.RFC3339),
// 	})
// }

// //#endregion

// //#region: Extracting and validating it

// // getTokenFromHeader extracts the JWT token from the request headers.
// // If everything occurrs well, returns the extracted jwt, otherwise, throw an
// // error.
// func (md *JWTMiddleware) getTokenFromHeader(c *gin.Context) (string, error) {
// 	authHeader := c.Request.Header.Get("Authorization")

// 	if authHeader == "" {
// 		return "", ErrMissingAuthorization
// 	}

// 	bearerHeader := strings.SplitN(authHeader, " ", 2)

// 	if bearerHeader[0] != "Bearer" {
// 		return "", ErrInvalidBearerHeader
// 	}

// 	return bearerHeader[1], nil
// }

// // getTokenFromCookie extracts the JWT token from the request cookies.
// // If everything occurrs well, returns the extracted jwt, otherwise, throw an
// // error.
// func (md *JWTMiddleware) getTokenFromCookie(c *gin.Context) (string, error) {
// 	return c.Cookie("session")
// }

// // getJWTClaims extract the JWT clains object
// func (md *JWTMiddleware) getJWTClaims(tokenString string) (*jwtClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return md.Key, nil
// 		})

// 	if err != nil || !token.Valid {
// 		return nil, ErrInvalidJWTToken
// 	}

// 	claims, ok := token.Claims.(*jwtClaims)

// 	if !ok {
// 		return nil, ErrInvalidJWTClaims
// 	}

// 	return claims, nil
// }

// // findSession get the current JWT session. This will search for token in the
// // request headers, if didn't find it, it'll use the cookies instead.
// // If no token was provided, or if fails to get the session, will return a
// // blank cookie and an error message, otherwise, it'll returns the session and
// // nil.
// //
// // See [getTokenFromHeader], [getTokenFromCookie], [getJWTClaims]
// func (md *JWTMiddleware) findSession(c *gin.Context) (Session, error) {
// 	tokeString, err := md.getTokenFromHeader(c)

// 	if err != nil {
// 		tokeString, err = md.getTokenFromCookie(c)

// 		if err != nil {
// 			return BLANK_SESSION, err
// 		}
// 	}

// 	claims, err := md.getJWTClaims(tokeString)

// 	if err != nil {
// 		return BLANK_SESSION, err
// 	}

// 	return claims.Session, nil
// }

// //#endregion

// // Middleware is the function passed through gin setup
// func (md *JWTMiddleware) Middleware(c *gin.Context) {
// 	session, err := md.findSession(c)

// 	if err != nil {
// 		// c.Error(err)
// 		// c.AbortWithStatus(http.StatusUnauthorized)
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.JsonError{
// 			Status: http.StatusUnauthorized,
// 			Errors: []string{err.Error()},
// 		})
// 		return
// 	}

// 	c.Set(GIN_JWT_SESSION_KEY, session)
// 	c.Next()
// }

// // RefreshHandler get the current session and generates a new jwt token.
// // Note that this function doesn't care about unvalidate the old one.
// func (md *JWTMiddleware) RefreshHandler(c *gin.Context) {
// 	session, ok := GetSession(c)

// 	if !ok {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	md.buildAndResponseToken(c, session, config.App.ApiDomain)
// }

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
