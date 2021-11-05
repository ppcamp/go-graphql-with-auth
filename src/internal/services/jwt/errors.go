package jwt

import "errors"

var (
	ErrInvalidJWTToken      = errors.New("invalid jwt token")
	ErrMissingAuthorization = errors.New("missing authorization header")
	ErrInvalidBearerHeader  = errors.New("invalid bearer header")
	ErrInvalidJWTClaims     = errors.New("invalid jwtClaims")
)
