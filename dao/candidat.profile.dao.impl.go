package dao

import (
	"lawencon.com/bootest/model"
)

type CandidatProfileDaoImpl struct{}

func (CandidatProfileDaoImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer catchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
