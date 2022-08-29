package handler

import (
	"log"
	"myapp/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	admin := c.Get("admin").(bool)
	log.Printf("login with ---> %v: %v", username, admin)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(20 * time.Minute).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		log.Printf("login err ---> %v", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginRespone{
		Token: t,
	})
}
