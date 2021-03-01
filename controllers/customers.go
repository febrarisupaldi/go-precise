package controllers

import (
	"net/http"

	"github.com/febrarisupaldi/go-learning-api/models"
	"github.com/labstack/echo/v4"
)

func GetAllCountry(c echo.Context) error {
	result, err := models.GetAllCountry()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddCustomers(c echo.Context) error {
	name := c.FormValue("customer_name")
	address := c.FormValue("customer_address")
	contact := c.FormValue("customer_contact")

	result, err := models.AddCustomers(name, address, contact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
