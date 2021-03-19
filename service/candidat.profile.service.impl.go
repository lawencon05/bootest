package service

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

type CandidatProfileServiceImpl struct{
	dao.CandidatProfileDao
}

func (candidateProfile CandidatProfileServiceImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer config.CatchError(&e)
	return candidateProfile.CandidatProfileDao.CreateCandidat(data)
}
