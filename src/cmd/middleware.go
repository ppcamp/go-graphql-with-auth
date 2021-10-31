package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-graphql-with-auth/internal/config"
)

var (
	ErrInvalidContentType = errors.New("invalid content type")
	ErrNotFound           = errors.New("not found")
	ErrMethodNotAllowed   = errors.New("method not allowed")
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

//#region: HTTP basic middlewares

// Define the allowed content type
func (a *Middleware) Options(c *gin.Context) {
	origin := c.GetHeader("Origin")

	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Max-Age", "3600")
	c.Header("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, OPTIONS, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Encoding, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")

	// c.Header("Content-Type", "application/json")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Max-Age", "3600")
		c.AbortWithStatus(http.StatusOK)
		return
	}

	if c.Request.Method == "POST" || c.Request.Method == "PUT" {
		isJson := strings.Contains(c.ContentType(), config.ContentTypeJSON)
		isFile := (c.Request.Method == "POST") && strings.Contains(c.ContentType(), config.ContentTypeMultipart)

		if !isJson && !isFile {
			c.Error(ErrInvalidContentType)
			c.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}
	}

	c.Next()
}

func (a *Middleware) Errors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors)
	}
}

func (a *Middleware) NotFound(c *gin.Context) {
	c.Error(ErrNotFound)
	c.AbortWithStatus(http.StatusNotFound)
}

func (a *Middleware) MethodNotAllowed(c *gin.Context) {
	c.Error(ErrMethodNotAllowed)
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

//#endregion

//#region: auth middleware

// BuildAuthMiddleware builds a method that will be used to validate the
// login configs. This method will use the auth service
// func (a *Middleware) BuildAuthMiddleware() *jwt.JWTMiddleware {
// 	authservice := auth.NewAuthService(a.Redis)

// 	return &jwt.JWTMiddleware{
// 		Key:              []byte(config.App.JWTSecret),
// 		Expires:          config.JWT_EXPIRATION_TIME,
// 		ApiAuthenticator: authservice.ApiAuthenticator,
// 	}
// }

//#endregion
