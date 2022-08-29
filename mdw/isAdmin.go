package mdw

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAdminMdv(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)
		log.Printf("-->IsAdminMdv %v", admin)
		if admin {
			next(c)
		}
		return echo.ErrUnauthorized
	}
}
