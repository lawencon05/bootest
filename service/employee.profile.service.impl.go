package service

import (
	"time"

	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var employeeDao dao.EmployeeProfileDao = dao.EmployeeProfileDaoImpl{}

type EmployeeProfileServiceImpl struct{}

func (EmployeeProfileServiceImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer config.CatchError(&e)
	*data.CreatedDate = model.Timestamp(time.Now())
	return employeeDao.CreateEmployee(data)
}
