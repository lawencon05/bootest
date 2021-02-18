package model

import "time"

type CandidateProfiles struct {
	BaseModel
	UserId   string    `gorm:"type:varchar(50)"`
	Users    Users     `gorm:"foreignKey:UserId"`
	Fullname string    `json:"fullname"`
	Address  string    `json:"address"`
	Dob      time.Time `json:"dob"`
}

func (CandidateProfiles) TableName() string {
	return "tb_m_cndt_profiles"
}
