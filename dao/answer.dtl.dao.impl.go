package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDtlDaoImpl struct{}

func (AnswerDtlDaoImpl) CreateAnswerDtl(data *models.AnswerDtl, tx *gorm.DB) error{	
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
