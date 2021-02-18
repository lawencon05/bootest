package service

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var employeeDao dao.EmployeeProfileDao = dao.EmployeeProfileDaoImpl{}

type EmployeeProfileServiceImpl struct{}

func (EmployeeProfileServiceImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer config.CatchError(&e)
	return employeeDao.CreateEmployee(data)
}
