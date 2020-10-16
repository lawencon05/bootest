package models

import "time"

type AnswerDtl struct {
	id          string
	answerHdrId string
	questionId  string
	answers     string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
