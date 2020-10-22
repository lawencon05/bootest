package services

import (
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/models"
)

type AnswerDtlService interface {
	CreateAnswerDtl(dtl *models.AnswerDtl, tx *gorm.DB) error
}
