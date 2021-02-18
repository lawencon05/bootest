package service

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var candidatDao dao.CandidatProfileDao = dao.CandidatProfileDaoImpl{}

type CandidatProfileServiceImpl struct{}

func (CandidatProfileServiceImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer config.CatchError(&e)
	return candidatDao.CreateCandidat(data)
}
