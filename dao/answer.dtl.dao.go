package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDtlDao interface {
	CreateAnswerDtl(data *models.AnswerDtl, tx *gorm.DB) error
}
