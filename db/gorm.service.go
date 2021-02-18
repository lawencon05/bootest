package db

import "gorm.io/gorm"

type GormService interface {
	Conn() (*gorm.DB, error)
}
