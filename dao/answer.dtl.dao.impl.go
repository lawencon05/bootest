package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type AnswerDtlDaoImpl struct{
	*gorm.DB
}

func (answerDtlDao AnswerDtlDaoImpl) CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
