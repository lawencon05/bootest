package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

type AnswerServiceImpl struct{
	dao.AnswerDao
	dao.AnswerDtlDao
	*gorm.DB
}

func (answerService AnswerServiceImpl) CreateAnswer(hdr *model.AnswerHdr, dtl *[]model.AnswerDtl) (e error) {
	defer config.CatchError(&e)
	return answerService.DB.Transaction(func(tx *gorm.DB) error {
		if err := answerService.CreateAnswerHdr(hdr, tx); err != nil {
			return err
		}

		for _, v := range *dtl {
			v.AnswerHdrId = *&hdr.Id
			if err := answerService.AnswerDtlDao.CreateAnswerDtl(&v, tx); err != nil {
				return err
			}
		}

		return nil
	})
}
