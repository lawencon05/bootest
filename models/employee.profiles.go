package models

import "time"

type EmployeeProfiles struct {
	id          string
	userId      string
	fullname    string
	nip         string
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
