package dao

import (
	"fmt"

	"lawencon.com/imamfarisi/models"
)

type UserDaoImpl struct{}

func (UserDaoImpl) CreateUser(user *models.Users) {
	result := newConn().Create(user)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	}	
}
