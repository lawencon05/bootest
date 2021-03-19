package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

type UserServiceImpl struct{
	dao.UserDao
}

func (userService UserServiceImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	result, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), 4)
	if err == nil {
		user.Pwd = string(result)
		return userService.UserDao.CreateUser(user)
	}
	return nil, err
}

func (userService UserServiceImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)
	u, err := userService.UserDao.GetUserById(id)
	u.Pwd = "" //do not show password !!!
	return u, err
}

func (userService UserServiceImpl) Login(username string, pwd string) (u model.Users, e error) {
	defer config.CatchError(&e)
	result, err := userService.UserDao.GetUserByUsername(username)
	if err == nil && result.Count > 0 {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Pwd), []byte(pwd))
		if err == nil {
			result.Pwd = "" //do not show password !!!
			v, _ := config.GenerateToken(username)
			result.Token = v
			return result, nil
		}
	}
	return model.Users{}, errors.New("Username/Password incorect")
}
