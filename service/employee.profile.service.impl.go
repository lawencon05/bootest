package service

import (
	"time"

	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var employeeDao dao.EmployeeProfileDao = dao.EmployeeProfileDaoImpl{}

type EmployeeProfileServiceImpl struct{}

func (EmployeeProfileServiceImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer catchError(&e)
	data.CreatedDate = time.Now()
	return employeeDao.CreateEmployee(data)
}
