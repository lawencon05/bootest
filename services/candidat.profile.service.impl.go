package services

import (
	"time"

	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var candidatDao dao.CandidatProfileDao = dao.CandidatProfileDaoImpl{}

type CandidatProfileServiceImpl struct{}

func (CandidatProfileServiceImpl) CreateCandidat(data *models.CandidateProfiles) error {
	data.CreatedDate = time.Now()
	return candidatDao.CreateCandidat(data)
}
