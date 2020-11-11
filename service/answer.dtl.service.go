package service

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/model"
)

type AnswerDtlService interface {
	CreateAnswerDtl(dtl *model.AnswerDtl, tx *gorm.DB) error
}
