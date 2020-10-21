package services

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var userDao dao.UserDao = dao.UserDaoImpl{}

type UserServiceImpl struct{}

func (UserServiceImpl) CreateUser(user *models.Users) (*models.Users, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), 4)
	if err == nil {
		user.Pwd = string(result)
		return userDao.CreateUser(user)
	}
	return nil, err
}

func (UserServiceImpl) GetUserById(id string) (models.Users, error) {
	return userDao.GetUserById(id)
}

func (UserServiceImpl) Login(username string, pwd string) (models.Users, error) {
	result, err := userDao.GetUserByUsername(username)
	fmt.Println(result, result.Count)
	if err == nil && result.Count > 0 {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Pwd), []byte(pwd))
		fmt.Println(err)
		if err == nil {
			return result, nil
		}
	}

	return models.Users{}, errors.New("Username/password not found")
}
