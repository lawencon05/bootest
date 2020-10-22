package services

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var dtlDao dao.AnswerDtlDao = dao.AnswerDtlDaoImpl{}

type AnswerDtlServiceImpl struct{}

func (AnswerDtlServiceImpl) CreateAnswerDtl(data *models.AnswerDtl, tx *gorm.DB) error {
	return dtlDao.CreateAnswerDtl(data, tx)
}
