package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var dtlDao dao.AnswerDtlDao = dao.AnswerDtlDaoImpl{}

type AnswerDtlServiceImpl struct{}

func (AnswerDtlServiceImpl) CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) (e error) {
	defer catchError(&e)
	return dtlDao.CreateAnswerDtl(data, tx)
}
