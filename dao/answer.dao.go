package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDao interface {
	CreateAnswerHdr(hdr *models.AnswerHdr, tx *gorm.DB) error
}
