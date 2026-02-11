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

func Transfer(c echo.Context) error {
	var (
		ctn       = c.Get(container.ContainerDefName).(di.Container)
		walletApp = ctn.Get(container.WalletApplicationDefName).(*application.Wallet)
		input     = dto.Transfer{}
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

	result, err := walletApp.Transfer(ctx, input)
	if err != nil {
		return response.DisplayCustomError(c, err)
	}

	return response.Ok(c, "successfully transfer", result)

}

func TransactionHistory(c echo.Context) error {
	var (
		ctn       = c.Get(container.ContainerDefName).(di.Container)
		walletApp = ctn.Get(container.WalletApplicationDefName).(*application.Wallet)
		ctx       = context.Background()
		userId    = c.Get("user_id")
	)

	ctx = context.WithValue(ctx, "user_id", userId)

	result, err := walletApp.GetTransactionHistory(ctx)
	if err != nil {
		return response.DisplayCustomError(c, err)
	}

	return response.Ok(c, "successfully transfer", result)
}

func TransactionRoutes(api *echo.Group, ctn *di.Container) *echo.Group {

	trx := api.Group("/transaction")
	trx.Use(middleware.AuthenticationMiddleware)
	trx.POST("/transfer", Transfer)
	trx.GET("/history", TransactionHistory)

	return trx
}
