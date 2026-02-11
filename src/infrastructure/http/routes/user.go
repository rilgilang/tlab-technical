package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"tlab/bootstrap/container"
	"tlab/src/application"
	"tlab/src/domain/sharedkernel/response"
	"tlab/src/infrastructure/http/middleware"
)

func Profile(c echo.Context) error {
	var (
		ctn     = c.Get(container.ContainerDefName).(di.Container)
		userApp = ctn.Get(container.UserApplicationDefName).(*application.User)
		ctx     = context.Background()
		userId  = c.Get("user_id")
	)

	ctx = context.WithValue(ctx, "user_id", userId)
	profile, err := userApp.GetProfile(ctx)
	if err != nil {
		return response.DisplayCustomError(c, err)
	}

	return response.Ok(c, "successfully get profile", profile)
}

func UserRoutes(api *echo.Group, ctn *di.Container) *echo.Group {

	user := api.Group("/users")
	user.Use(middleware.AuthenticationMiddleware)
	user.GET("/profile", Profile)

	return user
}
