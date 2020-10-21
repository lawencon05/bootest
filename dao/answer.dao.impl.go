package dao

import (
	"time"

	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDaoImpl struct{}

func (AnswerDaoImpl) CreateAnswer(hdr *models.AnswerHdr, dtl *[]models.AnswerDtl) error{
	return g.Transaction(func(tx *gorm.DB) error {
		hdr.CreatedDate = time.Now()
		if err := tx.Create(hdr).Error; err != nil {
			return err
		}

		for _, v := range *dtl {
			v.AnswerHdrId = *&hdr.Id
			v.CreatedDate = time.Now()
			if err := tx.Create(v).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
