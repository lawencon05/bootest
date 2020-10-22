package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDaoImpl struct{}

func (AnswerDaoImpl) CreateAnswerHdr(hdr *models.AnswerHdr, tx *gorm.DB) error {
	if err := tx.Create(hdr).Error; err != nil {
		return err
	}
	return nil
}
