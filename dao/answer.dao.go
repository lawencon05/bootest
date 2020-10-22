package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDao interface {
	CreateAnswerHdr(data *models.AnswerHdr, tx *gorm.DB) error
}
