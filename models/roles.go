package models

import "time"

type Roles struct {
	id          string
	roleName    string
	code        string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
