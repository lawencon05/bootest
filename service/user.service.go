package service

import (
	"lawencon.com/bootest/model"
)

type UserService interface {
	CreateUser(user *model.Users) (*model.Users, error)
	GetUserById(id string) (model.Users, error)
	Login(username string, pwd string) (model.Users, error)
}
