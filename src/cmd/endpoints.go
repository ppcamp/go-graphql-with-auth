package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-graphql-with-auth/internal/config"
	"github.com/ppcamp/go-graphql-with-auth/internal/controllers/app"
	"github.com/ppcamp/go-graphql-with-auth/internal/controllers/user"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/graphql"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
)

func SetupEngine(storage postgres.Storage) *gin.Engine {
	router := gin.New()

	// middlewares
	registerMiddlewares(router)

	// handlers
	schema := graphql.NewSchemaManager()
	userController := user.NewUserControllerBuilder(storage)
	appController := app.NewAppController(storage)

	// Endpoints unprotected
	schema.RegisterQuery("users", userController.QueryUsers())
	schema.RegisterMutation("createUser", userController.CreateUser())

	// Endpoints protected
	schema.RegisterAuthenticatedQuery("app", appController.QueryAppStatus())
	schema.RegisterAuthenticatedMutation("updateUser", userController.EditUser())

	// register
	router.Any("/graphql", schema.Handler())

	return router
}

func registerMiddlewares(router *gin.Engine) {
	middleware := NewMiddleware()

	// Return 500 on panic
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Response as JSON every c.Error in requests handler
	router.Use(middleware.Errors)

	// Handle OPTIONS and set default headers like CORS and Content-Type
	router.Use(middleware.Options)
	router.NoRoute(middleware.NotFound)
	router.NoMethod(middleware.MethodNotAllowed)

	// register a middleware to get all JWT auth
	authMiddleware := jwt.NewJwtMiddleware(
		config.App.JWTExp, []byte(config.App.JWTSecret))
	router.Use(authMiddleware.Middleware)
}
