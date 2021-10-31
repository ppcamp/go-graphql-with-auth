package jwt

import "errors"

var (
	ErrInvalidLoginPayload  = errors.New("invalid login payload")
	ErrInvalidJWTToken      = errors.New("invalid jwt token")
	ErrMissingAuthorization = errors.New("missing authorization header")
	ErrInvalidBearerHeader  = errors.New("invalid bearer header")
	ErrInvalidJWTClaims     = errors.New("invalid jwtClaims")
	ErrBlankAccessKey       = errors.New("invalid access key")
	ErrBlankSecretKey       = errors.New("invalid secret key")
)
