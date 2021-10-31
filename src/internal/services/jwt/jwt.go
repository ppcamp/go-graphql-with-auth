package jwt

import (
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

// Is the name of the key for the gin. Using this key we can get the value later
// on the gin context
const GIN_JWT_SESSION_KEY = "jwt_session"

// BlankSession is the default value returned when occurrs a problem to login
var BLANK_SESSION = Session{}

type LoginSession struct {
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Session struct {
	LoginSession // fields to put into jwt key

	Authenticated bool `json:"-"` // doesn't have any value
}

type UserSession struct {
	PhoneNumber string `json:"phone_number,omitempty"`
}

// jwtClaims is used by jwt middleware
type jwtClaims struct {
	Session
	jwt.StandardClaims
}

// GetSession get the current session of the gin.
// If there's no session in the req, returns false
func GetSession(c *gin.Context) (Session, bool) {
	session, ok := c.Get(GIN_JWT_SESSION_KEY)

	if !ok {
		return Session{}, false
	}

	ses := session.(Session)
	ses.Authenticated = true

	return ses, true
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
