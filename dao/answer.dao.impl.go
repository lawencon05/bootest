package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type AnswerDaoImpl struct{
	*gorm.DB
}

func (answerDao AnswerDaoImpl) CreateAnswerHdr(data *model.AnswerHdr, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
