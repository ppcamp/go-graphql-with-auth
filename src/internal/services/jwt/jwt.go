package jwt

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"github.com/ppcamp/go-graphql-with-auth/internal/config"
	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
)

// BlankSession is the default value returned when occurrs a problem to login
var BLANK_SESSION = Session{}

type LoginSession struct {
	UserId int64 `json:"userId,omitempty"`
}

type Session struct {
	*LoginSession // fields to put into jwt key

	Authenticated bool `json:"-"` // doesn't have any value
}

// jwtClaims is used by jwt middleware
type jwtClaims struct {
	Session
	jwt.StandardClaims
}

// generateJwt generates a token to the current session for a given amount
// "expiration" of time. This also will set the "claims", i.e, the object
// that will be encoded into JWT
func generateJwt(session Session, expiration time.Time) *jwt.Token {
	claims := jwtClaims{
		Session: session,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// generateJwtAsSignedString does the same as [generateJwt], however, it'll
// return the jwt as string
func generateJwtAsSignedString(session Session, expiration time.Time, key []byte) (string, error) {
	token := generateJwt(session, expiration)
	signedString, err := token.SignedString(key)
	return signedString, err
}

// buildAndResponseToken generates a signed endpoint with expiration
func buildAndResponseToken(
	c *gin.Context,
	session Session,
	domain string,
	expires time.Duration,
	key []byte,
) (loginmodels.TokenResponse, error) {
	exp := time.Now().UTC().Add(expires)
	signedString, err := generateJwtAsSignedString(session, exp, key)

	if err != nil {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return loginmodels.TokenResponse{}, err
	}

	c.SetCookie("session", signedString, int(expires.Seconds()), "/", domain, false, false)
	return loginmodels.TokenResponse{
		Token:   signedString,
		Expires: exp.Format(time.RFC3339),
	}, nil
}

func BuildToken(ctx context.Context, s Session) (loginmodels.TokenResponse, error) {
	c := ctx.(*gin.Context)

	return buildAndResponseToken(
		c,
		s,
		config.App.ApiDomain,
		config.App.JWTExp,
		[]byte(config.App.JWTSecret))
}

// GetSession get the current session registered into context.
func GetSession(c context.Context) (Session, error) {
	session := c.Value(GIN_JWT_SESSION_KEY)

	// logrus.WithField("sess", session).Debug("getSession")

	if ses, ok := session.(Session); ok {
		if ses.LoginSession == nil {
			return Session{}, ErrMissingAuthorization
		}
		ses.Authenticated = true
		return ses, nil
	} else {
		// logrus.WithField("cox", session).Debug(strings.Repeat("x", 50))
		return Session{}, ErrMissingAuthorization
	}
}
