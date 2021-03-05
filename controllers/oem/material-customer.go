package oem

import (
	"net/http"
	"github.com/febrarisupaldi/go-precise/models/oem"
	"github.com/labstack/echo/v4"
)



func AllMaterialCustomer(c echo.Context) error{
	result, err := oem.AllMaterialCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}