package service

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

type EmployeeProfileServiceImpl struct{
	dao.EmployeeProfileDao
}

func (empProfile EmployeeProfileServiceImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer config.CatchError(&e)
	return empProfile.EmployeeProfileDao.CreateEmployee(data)
}
