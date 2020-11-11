package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/model"
)

type AnswerDaoImpl struct{}

func (AnswerDaoImpl) CreateAnswerHdr(data *model.AnswerHdr, tx *gorm.DB) (e error) {
	defer catchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
