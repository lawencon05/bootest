package model

type Users struct {
	Email  string `json:"email" gorm:"unique"`
	Pwd    string `json:"pwd"`
	RoleId string `json:"roleId" gorm:"column:role_id;type:varchar(50)"`
	Roles  Roles  `gorm:"foreignKey:RoleId"`
	Token  string `json:"token" gorm:"-"`
	BaseModel
}

func (Users) TableName() string {
	return "tb_m_users"
}
