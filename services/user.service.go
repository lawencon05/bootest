package services

import (
	"lawencon.com/imamfarisi/models"
)

type UserService interface {
	CreateUser(user *models.Users) (*models.Users, error)
	GetUserById(id string) (models.Users, error)
	Login(username string, pwd string) (models.Users, error)
}
