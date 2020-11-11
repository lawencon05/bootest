package service

import (
	"lawencon.com/bootest/model"
)

type AnswerService interface {
	CreateAnswer(hdr *model.AnswerHdr, dtl *[]model.AnswerDtl) error
}
