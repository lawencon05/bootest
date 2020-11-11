package controller

import (
	"github.com/labstack/echo"

	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var employeeProfileService service.EmployeeProfileService = service.EmployeeProfileServiceImpl{}

func SetEmployeeProfile(c *echo.Group) {
	c.POST("/employee", createEmployeeProfile)
}

func createEmployeeProfile(c echo.Context) (e error) {
	defer catchError(&e)
	data := new(model.EmployeeProfiles)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	var err = employeeProfileService.CreateEmployee(data)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
