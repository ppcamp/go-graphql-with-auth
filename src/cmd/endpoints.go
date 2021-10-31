package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ppcamp/go-graphql-with-auth/internal/controllers/app"
	"github.com/ppcamp/go-graphql-with-auth/internal/controllers/user"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/graphql"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
)

func SetupEngine(storage postgres.Storage) *gin.Engine {
	router := gin.New()

	registerMiddlewares(router)

	// handlers
	schemaManager := graphql.NewSchemaManager()
	userController := user.NewUserControllerBuilder(storage)
	appController := app.NewAppController(storage)

	// Endpoints
	schemaManager.RegisterQuery("app", appController.QueryAppStatus())
	schemaManager.RegisterQuery("user", userController.QueryHello())
	schemaManager.RegisterMutation("createUser", userController.CreateUser())

	gql := router.Group("/graphql")
	{
		gql.POST("", schemaManager.Handler())
		gql.GET("", schemaManager.Handler())
	}
	return router
}

func registerMiddlewares(router *gin.Engine) {
	middleware := NewMiddleware()

	// router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/appstatus"))
	// router.GET("/appstatus", appcontroller.HandleAppStatus)

	// Return 500 on panic
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Response as JSON every c.Error in requests handler
	router.Use(middleware.Errors)

	// Handle OPTIONS and set default headers like CORS and Content-Type
	router.Use(middleware.Options)
	router.NoRoute(middleware.NotFound)
	router.NoMethod(middleware.MethodNotAllowed)

}
