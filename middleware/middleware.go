package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey : []byte("D53A4C62525D45AF9C0FD6013A0084655FD7D52A"),
})