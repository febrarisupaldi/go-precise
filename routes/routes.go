package routes

import (

	"github.com/febrarisupaldi/go-precise/controllers"
	mw "github.com/febrarisupaldi/go-precise/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", controllers.Login)
	
	m := e.Group("/master")
	m.GET("/country", controllers.AllCountry, mw.IsAuthenticated)
	m.GET("/country/check", controllers.CheckCountry, mw.IsAuthenticated)
	m.GET("/country/:id", controllers.ShowCountry, mw.IsAuthenticated)
	m.POST("/country", controllers.AddCountry, mw.IsAuthenticated)
	m.PUT("/country/:id", controllers.UpdateCountry, mw.IsAuthenticated)
	m.DELETE("/country/:id", controllers.DeleteCountry, mw.IsAuthenticated)
	
	//e.POST("/generate", controllers.GenerateHashPassword)
	return e
}
