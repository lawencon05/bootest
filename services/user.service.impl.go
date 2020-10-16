package services

import (
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var userDao dao.UserDao = dao.UserDaoImpl{}

type UserServiceImpl struct{}

func (UserServiceImpl) CreateUser(user *models.Users) {
	userDao.CreateUser(user)
}
