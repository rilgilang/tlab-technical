package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"strings"
	"tlab/bootstrap/config"
	"tlab/bootstrap/container"
	"tlab/src/domain/sharedkernel/response"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctn = c.Get(container.ContainerDefName).(di.Container)
			cfg = ctn.Get(container.ConfigDefName).(config.Config)
		)

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "missing authorization header")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.Unauthorized(c, "Invalid Authorization header format")
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil {
			return response.Unauthorized(c, "invalid token")
		}

		if !token.Valid {
			return response.Unauthorized(c, "token is not valid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return response.DisplayCustomError(c, errors.New("server_error"))
		}

		userID, ok := claims["id"].(string)
		if !ok {
			return response.DisplayCustomError(c, errors.New("server_error"))
		}
		c.Set("user_id", userID)

		// Continue to the next handler
		return next(c)
	}
}
