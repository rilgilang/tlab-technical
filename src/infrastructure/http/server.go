package http

import (
	"tlab/src/infrastructure/http/middleware"
	"tlab/src/infrastructure/http/routes"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
)

func SetupRoutes(
	e *echo.Echo,
	container *di.Container,
) {

	api := e.Group("/api")
	//api versioning
	//v1 := api.Group("/v1")

	// Inject container to echo context
	e.Use(middleware.InjectContainer(*container))

	//public routes
	routes.AuthRoutes(api, container)
	routes.TransactionRoutes(api, container)
}
