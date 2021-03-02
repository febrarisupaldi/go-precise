package routes

import (

	"github.com/febrarisupaldi/go-precise/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/country", controllers.GetAllCountry)
	e.POST("/country", controllers.AddCountry)
	e.PUT("/country/:id", controllers.UpdateCountry)
	e.DELETE("/country/:id", controllers.DeleteCountry)
	e.POST("/login", controllers.Login)
	//e.POST("/generate", controllers.GenerateHashPassword)
	return e
}
