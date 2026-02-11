package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"tlab/bootstrap/container"
	"tlab/src/application"
	"tlab/src/domain/sharedkernel/response"
	"tlab/src/infrastructure/http/dto"
	"tlab/src/infrastructure/http/middleware"
)

func TopUpBalance(c echo.Context) error {
	var (
		ctn       = c.Get(container.ContainerDefName).(di.Container)
		walletApp = ctn.Get(container.WalletApplicationDefName).(*application.Wallet)
		input     = dto.TopUp{}
		ctx       = context.Background()
		userId    = c.Get("user_id")
	)

	ctx = context.WithValue(ctx, "user_id", userId)
	if err := c.Bind(&input); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err := c.Validate(&input); err != nil {
		return response.ValidationError(c, err)
	}

	result, err := walletApp.TopUpBalance(ctx, input)
	if err != nil {
		return response.Error(c, 500, "internal server error", nil)
	}

	//return response.Ok(c, "successfully charge", trx)
	return response.Ok(c, "successfully top up balance", result)

}

func GetBalance(c echo.Context) error {
	var (
		ctn       = c.Get(container.ContainerDefName).(di.Container)
		walletApp = ctn.Get(container.WalletApplicationDefName).(*application.Wallet)
		ctx       = context.Background()
		userId    = c.Get("user_id")
	)

	ctx = context.WithValue(ctx, "user_id", userId)

	result, err := walletApp.GetBalance(ctx)
	if err != nil {
		return response.Error(c, 500, "internal server error", nil)
	}

	//return response.Ok(c, "successfully charge", trx)
	return response.Ok(c, "successfully get balance", result)

}

func WalletRoutes(api *echo.Group, ctn *di.Container) *echo.Group {

	wallet := api.Group("/wallet")
	wallet.Use(middleware.AuthenticationMiddleware)
	wallet.POST("/topup", TopUpBalance)
	wallet.GET("/balance", GetBalance)

	return wallet
}
