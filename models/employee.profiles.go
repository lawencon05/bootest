package models

type EmployeeProfiles struct {
	BaseModels
	UserId   string `json:"userId"`
	Users    Users  `gorm:"foreignKey:UserId"`
	Fullname string `json:"fullname"`
	Nip      string `json:"nip"`
}

func (EmployeeProfiles) TableName() string {
	return "tb_m_emp_profiles"
}
