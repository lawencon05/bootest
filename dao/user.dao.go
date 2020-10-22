package dao

import (
	"lawencon.com/imamfarisi/models"
)

type UserDao interface {
	CreateUser(data *models.Users) (*models.Users, error)
	GetUserById(id string) (models.Users, error)
	GetUserByUsername(username string) (models.Users, error)
}
