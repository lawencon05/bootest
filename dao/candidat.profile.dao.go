package dao

import (
	"lawencon.com/bootest/model"
)

type CandidatProfileDao interface {
	CreateCandidat(data *model.CandidateProfiles) error
}
