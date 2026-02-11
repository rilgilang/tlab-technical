package dto

import "github.com/labstack/echo/v4"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Detail  interface{} `json:"meta"`
}

func JsonResponse(c echo.Context, code int, data interface{}, message, status string, detail interface{}) (err error) {
	return c.JSON(code, &Response{
		Status:  status,
		Message: message,
		Data:    data,
		Detail:  detail,
	})
}
