package dao

import (
	"lawencon.com/imamfarisi/models"
)

type EmployeeProfileDaoImpl struct{}

func (EmployeeProfileDaoImpl) CreateEmployee(data *models.EmployeeProfiles) error {
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
