package services

import (
	"lawencon.com/imamfarisi/models"
)

type EmployeeProfileService interface {
	CreateEmployee(data *models.EmployeeProfiles) error
}
