package models

type Users struct {
	Email  string `json:"email"`
	Pwd    string `json:"pwd"`
	RoleId string `json:"roleId"`
	Token  string `json:"token" gorm:"-" sql:"-"`
	BaseModels
}

func (Users) TableName() string {
	return "tb_m_users"
}
