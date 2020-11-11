package service

import (
	"lawencon.com/bootest/model"
)

type CandidatProfileService interface {
	CreateCandidat(data *model.CandidateProfiles) error
}
