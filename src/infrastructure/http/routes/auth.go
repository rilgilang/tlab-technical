package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"tlab/bootstrap/container"
	"tlab/src/application"
	"tlab/src/domain/sharedkernel/response"
	"tlab/src/infrastructure/http/dto"
)

func Login(c echo.Context) error {
	var (
		ctn     = c.Get(container.ContainerDefName).(di.Container)
		userApp = ctn.Get(container.UserApplicationDefName).(*application.User)
		input   = dto.LoginInput{}
		ctx     = c.Request().Context()
	)

	if err := c.Bind(&input); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err := c.Validate(&input); err != nil {
		return response.ValidationError(c, err)
	}

	token, err := userApp.Login(ctx, input)
	if err != nil {
		return response.DisplayCustomError(c, err)
	}

	return response.Ok(c, "successfully login", token)
}

func Register(c echo.Context) error {
	var (
		ctn     = c.Get(container.ContainerDefName).(di.Container)
		userApp = ctn.Get(container.UserApplicationDefName).(*application.User)
		input   = dto.RegisterInput{}
		ctx     = c.Request().Context()
	)

	if err := c.Bind(&input); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err := c.Validate(&input); err != nil {
		return response.ValidationError(c, err)
	}

	err := userApp.Register(ctx, input)
	if err != nil {
		return response.DisplayCustomError(c, err)
	}

	return response.Ok(c, "successfully register", nil)
}

func AuthRoutes(api *echo.Group, ctn *di.Container) *echo.Group {

	auth := api.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/register", Register)

	return auth
}
