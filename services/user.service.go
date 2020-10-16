package services

import "lawencon.com/imamfarisi/models"

type UserService interface {
	CreateUser(user *models.Users)
}
