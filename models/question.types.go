package models

import "time"

type QuestionTypes struct {
	id          string
	qstnName    string
	code        string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
