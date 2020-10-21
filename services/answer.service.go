package services

import (
	"lawencon.com/imamfarisi/models"
)

type AnswerService interface {
	CreateAnswer(hdr *models.AnswerHdr, dtl *[]models.AnswerDtl) error
}
