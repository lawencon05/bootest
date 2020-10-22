package services

import (
	"time"

	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var employeeDao dao.EmployeeProfileDao = dao.EmployeeProfileDaoImpl{}

type EmployeeProfileServiceImpl struct{}

func (EmployeeProfileServiceImpl) CreateEmployee(data *models.EmployeeProfiles) error {
	data.CreatedDate = time.Now()
	return employeeDao.CreateEmployee(data)
}
