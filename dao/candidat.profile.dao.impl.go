package dao

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type CandidatProfileDaoImpl struct{}

func (CandidatProfileDaoImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
