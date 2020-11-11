package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"lawencon.com/bootest/config"
)

func SetInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err.Error())
}

func catchError(e *error) {
	config.CatchError(e)
}
