package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func SetInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	result, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, string(result))
}

func resErr(c echo.Context, err error) error {
	result, err := json.Marshal(err.Error())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, string(result))
}
