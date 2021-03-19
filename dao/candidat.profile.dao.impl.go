package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
)

type CandidatProfileDaoImpl struct{
	*gorm.DB
}

func (candidateProfileDao CandidatProfileDaoImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer config.CatchError(&e)
	result := candidateProfileDao.DB.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
