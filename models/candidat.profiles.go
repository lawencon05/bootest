package models

import "time"

type CandidateProfiles struct {
	BaseModels
	UserId   string
	Users    Users     `gorm:"foreignKey:UserId"`
	Fullname string    `json:"fullname"`
	Address  string    `json:"address"`
	Dob      time.Time `json:"dob"`
}

func (CandidateProfiles) TableName() string {
	return "tb_m_cndt_profiles"
}
