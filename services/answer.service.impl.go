package services

import (
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var answerDao dao.AnswerDao = dao.AnswerDaoImpl{}

type AnswerServiceImpl struct{}

func (AnswerServiceImpl) CreateAnswer(hdr *models.AnswerHdr, dtl *[]models.AnswerDtl) error {
	return answerDao.CreateAnswer(hdr, dtl)
}
