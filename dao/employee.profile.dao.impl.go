package dao

import (
	"lawencon.com/bootest/model"
)

type EmployeeProfileDaoImpl struct{}

func (EmployeeProfileDaoImpl) CreateEmployee(data *model.EmployeeProfiles) (e error) {
	defer catchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
