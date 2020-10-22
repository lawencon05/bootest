package dao

import (
	"lawencon.com/imamfarisi/models"
)

type CandidatProfileDao interface {
	CreateCandidat(data *models.CandidateProfiles) error
}
