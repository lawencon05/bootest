package models

import "time"

type Questions struct {
	id          string
	qstn        string
	qstnType    string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
