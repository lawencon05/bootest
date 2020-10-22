package models

type Users struct {
	Email  string `json:"email" gorm:"unique"`
	Pwd    string `json:"pwd"`
	RoleId string `json:"roleId"`
	Token  string `json:"token" gorm:"-"`
	BaseModels
}

func (Users) TableName() string {
	return "tb_m_users"
}
