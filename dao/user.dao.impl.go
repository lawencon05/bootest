package dao

import (
	"lawencon.com/imamfarisi/models"
)

type UserDaoImpl struct{}

func (UserDaoImpl) CreateUser(user *models.Users) (*models.Users, error){
	result := g.Create(user)
	if result.Error == nil {
		return user, nil	
	}
	return nil, result.Error
}

func (UserDaoImpl) GetUserById(id string) (models.Users, error) {
	data := models.Users{BaseModels: models.BaseModels{Id: id}}
	result := g.First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil	
	}
	return data, result.Error
}

func (UserDaoImpl) GetUserByUsername(username string) (models.Users, error) {
	var data models.Users
	result := g.Where("email = ?", username).Find(&data)
	
	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil	
	}
	return data, result.Error
} 
