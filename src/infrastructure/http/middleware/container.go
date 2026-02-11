package middleware

import (
	container2 "tlab/bootstrap/container"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
)

func InjectContainer(container di.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(string(container2.ContainerDefName), container)
			return next(c)
		}
	}
}
