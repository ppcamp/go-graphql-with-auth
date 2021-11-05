package jwt

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

// Is the name of the key for the gin. Using this key we can get the value later
// on the gin context
const GIN_JWT_SESSION_KEY = "jwt_session"

type JWTMiddleware struct {
	Expires time.Duration
	Key     []byte
}

func NewJwtMiddleware(exp time.Duration, secret []byte) *JWTMiddleware {
	return &JWTMiddleware{
		Expires: exp,
		Key:     secret,
	}
}

//#region: Extracting and validating it

// getTokenFromHeader extracts the JWT token from the request headers.
//
// If everything occurrs well, returns the extracted jwt, otherwise, throw an
// error.
func (md *JWTMiddleware) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return "", ErrMissingAuthorization
	}

	bearerHeader := strings.SplitN(authHeader, " ", 2)

	if bearerHeader[0] != "Bearer" {
		return "", ErrInvalidBearerHeader
	}

	return bearerHeader[1], nil
}

// getTokenFromCookie extracts the JWT token from the request cookies.
//
// If everything occurrs well, returns the extracted jwt, otherwise, throw an
// error.
func (md *JWTMiddleware) getTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie("session")
}

// getJWTClaims extract the JWT clains object
func (md *JWTMiddleware) getJWTClaims(tokenString string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return md.Key, nil
		})

	if err != nil || !token.Valid {
		return nil, ErrInvalidJWTToken
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, ErrInvalidJWTClaims
	}

	return claims, nil
}

// findSession get the current JWT session.
//
// This will search for token in the
// request headers, if didn't find it, it'll use the cookies instead.
// If no token was provided, or if fails to get the session, will return a
// blank cookie and an error message, otherwise, it'll returns the session and
// nil.
//
// See [getTokenFromHeader], [getTokenFromCookie], [getJWTClaims]
func (md *JWTMiddleware) findSession(c *gin.Context) (Session, error) {
	tokeString, err := md.getTokenFromHeader(c)

	if err != nil {
		tokeString, err = md.getTokenFromCookie(c)

		if err != nil {
			return BLANK_SESSION, err
		}
	}

	claims, err := md.getJWTClaims(tokeString)

	if err != nil {
		return BLANK_SESSION, err
	}

	return claims.Session, nil
}

//#endregion

// Middleware is the function passed through gin setup.
//
// This function will register the header auth into context
func (md *JWTMiddleware) Middleware(c *gin.Context) {
	session, err := md.findSession(c)
	if err != nil {
		c.Set(GIN_JWT_SESSION_KEY, session)
	}
	c.Next()
}
