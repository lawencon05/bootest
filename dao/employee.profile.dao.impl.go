package dao

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type EmployeeProfileDaoImpl struct{}

func (EmployeeProfileDaoImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
