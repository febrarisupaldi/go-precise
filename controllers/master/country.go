package master

import (
	"strconv"
	"net/http"

	"github.com/febrarisupaldi/go-precise/models/master"
	"github.com/labstack/echo/v4"
)

func AllCountry(c echo.Context) error {
	result, err := master.AllCountry()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowCountry(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	result, err := master.ShowCountry(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddCountry(c echo.Context) error {
	code := c.FormValue("country_code")
	name := c.FormValue("country_name")
	by := c.FormValue("created_by")

	result, err := master.AddCountry(code, name, by)
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
	reason := c.FormValue("reason")

	result, err := master.UpdateCountry(id, code, name, by, reason)
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
	by := c.FormValue("deleted_by")
	reason := c.FormValue("reason")

	result, err := master.DeleteCountry(id, by, reason)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func CheckCountry(c echo.Context)error{
	tipe := c.QueryParam("type")
	value := c.QueryParam("value")

	result, err := master.CheckCountry(tipe, value)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]int{"status": http.StatusOK, "message":result})
}
