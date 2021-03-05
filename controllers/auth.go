package controllers

import (
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"github.com/febrarisupaldi/go-precise/models"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error{
	user_id := c.FormValue("user_id")
	password := c.FormValue("password")

	res, id, err := models.Login(user_id, password)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages":err.Error(),
		})
	}

	if !res{
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	t, err := token.SignedString([]byte("D53A4C62525D45AF9C0FD6013A0084655FD7D52A"))
	if err!= nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages":err.Error(),
		})
	}
	logId := strconv.FormatInt(id, 10)
	time := time.Now().Add(time.Hour * 8).String()
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":t,
		"expired_in": time,
		"log_id" : logId,
	})
}

// func GenerateHashPassword(c echo.Context) error{
// 	password := c.FormValue("password")
// 	hash, _ := models.HashPassword(password)
// 	return c.JSON(http.StatusOK, hash)
// }