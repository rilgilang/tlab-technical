package middleware

import (
	"net/http"
	"strings"
	"tlab/bootstrap/config"
	"tlab/bootstrap/container"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctn = c.Get(container.ContainerDefName).(di.Container)
			cfg = ctn.Get(container.ConfigDefName).(config.Config)
		)

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is not valid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error extracting claims")
		}

		userID, ok := claims["id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error extracting user ID from claims")
		}

		c.Set("user_id", userID)

		// Continue to the next handler
		return next(c)
	}
}
