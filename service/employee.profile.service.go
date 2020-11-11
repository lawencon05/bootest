package service

import (
	"lawencon.com/bootest/model"
)

type EmployeeProfileService interface {
	CreateEmployee(data *model.EmployeeProfiles) error
}
