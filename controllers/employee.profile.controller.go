package controllers

import (
	"net/http"

	"github.com/labstack/echo"

	"lawencon.com/imamfarisi/models"
	"lawencon.com/imamfarisi/services"
)

var employeeProfileService services.EmployeeProfileService = services.EmployeeProfileServiceImpl{}

func SetEmployeeProfile(c *echo.Group) {
	c.POST("/employee", createEmployeeProfile)
}

func createEmployeeProfile(c echo.Context) error {
	data := new(models.EmployeeProfiles)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var err = employeeProfileService.CreateEmployee(data)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
