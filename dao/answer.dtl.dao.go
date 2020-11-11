package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/model"
)

type AnswerDtlDao interface {
	CreateAnswerDtl(data *model.AnswerDtl, tx *gorm.DB) error
}
