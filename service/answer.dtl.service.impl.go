package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

type AnswerDtlServiceImpl struct{
	dao.AnswerDtlDao
	*gorm.DB
}

func (answerDtlService AnswerDtlServiceImpl) CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return answerDtlService.AnswerDtlDao.CreateAnswerDtl(data, tx)
}
