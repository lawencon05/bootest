package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/model"
)

type AnswerDao interface {
	CreateAnswerHdr(data *model.AnswerHdr, tx *gorm.DB) error
}
