package dao

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDtlDao interface {
	CreateAnswerDtl(dtl *models.AnswerDtl, tx *gorm.DB) error
}
