package models

import "time"

type CandidatProfiles struct {
	id          string
	userId      string
	employeeId  string
	fullname    string
	address     string
	dob         time.Time
	isActive    bool
	createdBy   string
	updatedBy   string
	createdDate time.Time
	updatedDate time.Time
}
