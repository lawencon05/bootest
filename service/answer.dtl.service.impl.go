package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var dtlDao dao.AnswerDtlDao = dao.AnswerDtlDaoImpl{}

type AnswerDtlServiceImpl struct{}

func (AnswerDtlServiceImpl) CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return dtlDao.CreateAnswerDtl(data, tx)
}
