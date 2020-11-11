package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/model"
)

type AnswerDtlDaoImpl struct{}

func (AnswerDtlDaoImpl) CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) (e error) {
	defer catchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
