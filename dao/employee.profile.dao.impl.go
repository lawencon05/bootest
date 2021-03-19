package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type EmployeeProfileDaoImpl struct{
	*gorm.DB
}

func (empProfileDao EmployeeProfileDaoImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer config.CatchError(&e)
	result := empProfileDao.DB.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
