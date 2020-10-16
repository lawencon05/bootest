package models

type Users struct {
	Id     string `json : "id"`
	Email  string `json : "email"`
	Pwd    string `json : "pwd"`
	RoleId string `json : roleId`
	BaseModels
}

type Tabler interface {
	TableName() string
}

func (Users) TableName() string {
	return "tb_m_users"
}
