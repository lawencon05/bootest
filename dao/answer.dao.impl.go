package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDaoImpl struct{}

func (AnswerDaoImpl) CreateAnswerHdr(data *models.AnswerHdr, tx *gorm.DB) error {
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
