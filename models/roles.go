package models

type Roles struct {
	BaseModels
	RoleName string `json:"name" gorm:"column:role_name"`
	RoleCode string `json:"code" gorm:"unique;column:code"`
}
