package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	e *echo.Echo
	eg *echo.Group
	db *gorm.DB
)

func SetInit(echoParam *echo.Echo, echoGroupParam *echo.Group, gormDbParam *gorm.DB) {
	e = echoParam
	eg = echoGroupParam
	db = gormDbParam
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})

	initControllers()
}

func initControllers() {
	setUser()
	setAnswer()
	setCandidatProfile()
	setEmployeeProfile()
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err.Error())
}
