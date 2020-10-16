package dao

import "lawencon.com/imamfarisi/models"

type UserDao interface {
	CreateUser(user *models.Users)
}
