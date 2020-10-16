package models

import "time"

type AnswerHdr struct {
	id          string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
