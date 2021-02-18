package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var answerDao dao.AnswerDao = dao.AnswerDaoImpl{}

type AnswerServiceImpl struct{}

func (AnswerServiceImpl) CreateAnswer(hdr *model.AnswerHdr, dtl *[]model.AnswerDtl) (e error) {
	defer config.CatchError(&e)
	return g.Transaction(func(tx *gorm.DB) error {
		if err := answerDao.CreateAnswerHdr(hdr, tx); err != nil {
			return err
		}

		for _, v := range *dtl {
			v.AnswerHdrId = *&hdr.Id
			if err := dtlDao.CreateAnswerDtl(&v, tx); err != nil {
				return err
			}
		}

		return nil
	})
}
