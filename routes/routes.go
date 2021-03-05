package routes

import (
	"github.com/febrarisupaldi/go-precise/controllers"
	"github.com/febrarisupaldi/go-precise/controllers/master"
	"github.com/febrarisupaldi/go-precise/controllers/oem"
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
	m.GET("/country", master.AllCountry, mw.IsAuthenticated)
	m.GET("/country/check", master.CheckCountry, mw.IsAuthenticated)
	m.GET("/country/:id", master.ShowCountry, mw.IsAuthenticated)
	m.POST("/country", master.AddCountry, mw.IsAuthenticated)
	m.PUT("/country/:id", master.UpdateCountry, mw.IsAuthenticated)
	m.DELETE("/country/:id", master.DeleteCountry, mw.IsAuthenticated)

	o:= e.Group("/oem")
	o.GET("/material-customer", oem.AllMaterialCustomer, mw.IsAuthenticated)
	
	//e.POST("/generate", controllers.GenerateHashPassword)
	return e
}
