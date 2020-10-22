package dao

import (
	"lawencon.com/imamfarisi/models"
)

type CandidatProfileDaoImpl struct{}

func (CandidatProfileDaoImpl) CreateCandidat(data *models.CandidateProfiles) error {
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
