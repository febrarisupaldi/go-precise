package controllers

import (
	"strconv"
	"net/http"

	"github.com/febrarisupaldi/go-precise/models"
	"github.com/labstack/echo/v4"
)

func GetAllCountry(c echo.Context) error {
	result, err := models.GetAllCountry()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddCountry(c echo.Context) error {
	code := c.FormValue("country_code")
	name := c.FormValue("country_name")
	by := c.FormValue("created_by")

	result, err := models.AddCountry(code, name, by)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateCountry(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	code := c.FormValue("country_code")
	name := c.FormValue("country_name")
	by := c.FormValue("updated_by")

	result, err := models.UpdateCountry(id, code, name, by)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)

}

func DeleteCountry(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCountry(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
