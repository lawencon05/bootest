package dao

import (
	"lawencon.com/imamfarisi/models"
)

type EmployeeProfileDao interface {
	CreateEmployee(data *models.EmployeeProfiles) error
}
