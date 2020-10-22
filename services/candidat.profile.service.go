package services

import (
	"lawencon.com/imamfarisi/models"
)

type CandidatProfileService interface {
	CreateCandidat(data *models.CandidateProfiles) error
}
