package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"tlab/src/domain/sharedkernel/response"
	"tlab/src/infrastructure/http/dto"
)

func Charge(c echo.Context) error {
	var (
		merchantId = c.Get("merchant_id").(string)
		//ctn            = c.Get(container.ContainerDefName).(di.Container)
		//transactionApp = ctn.Get(container.TransactionApplicationDefName).(*application.Transaction)
		input = dto.ChargeTransaction{}
		//ctx            = c.Request().Context()
	)

	if err := c.Bind(&input); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err := c.Validate(&input); err != nil {
		return response.ValidationError(c, err)
	}

	input.MerchantId = merchantId
	//trx, err := transactionApp.Charge(ctx, input)
	//if err != nil {
	//	return response.Error(c, 500, "internal server error", nil)
	//}

	//return response.Ok(c, "successfully charge", trx)
	return response.Ok(c, "successfully charge", "trx")

}

func TransactionRoutes(api *echo.Group, ctn *di.Container) *echo.Group {

	trx := api.Group("/transaction")
	//trx.Use(middleware.SecretKeyMiddleware(ctn))
	trx.POST("/charge", Charge)

	return trx
}
