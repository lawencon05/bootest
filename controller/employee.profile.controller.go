package controller

import (
	"github.com/labstack/echo/v4"

	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var employeeProfileService service.EmployeeProfileService

func setEmployeeProfile() {
	eg.POST("/employee", createEmployeeProfile)

	employeeProfileService = service.EmployeeProfileServiceImpl{
		EmployeeProfileDao: dao.EmployeeProfileDaoImpl{
			DB: db,
		},
	}
}

func createEmployeeProfile(c echo.Context) (e error) {
	defer config.CatchError(&e)
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
