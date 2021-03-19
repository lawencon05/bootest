package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type UserDaoImpl struct{
	*gorm.DB
}

func (userDao UserDaoImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	result := userDao.DB.Create(user)
	if result.Error == nil {
		return user, nil
	}
	return nil, result.Error
}

func (userDao UserDaoImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)
	data := model.Users{BaseModel: model.BaseModel{Id: id}}
	result := userDao.DB.First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

func (userDao UserDaoImpl) GetUserByUsername(username string) (u model.Users, e error) {
	defer config.CatchError(&e)
	var data model.Users
	result := userDao.DB.Where("email = ?", username).Find(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}
