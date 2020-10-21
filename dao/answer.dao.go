package dao

import (
	"lawencon.com/imamfarisi/models"
)

type AnswerDao interface {
	CreateAnswer(hdr *models.AnswerHdr, dtl *[]models.AnswerDtl) error
}
