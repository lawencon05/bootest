package model

type EmployeeProfiles struct {
	BaseModel
	UserId   string `json:"userId" gorm:"type:varchar(50)"`
	Users    Users  `gorm:"foreignKey:UserId"`
	Fullname string `json:"fullname"`
	Nip      string `json:"nip"`
}

func (EmployeeProfiles) TableName() string {
	return "tb_m_emp_profiles"
}
