package dao

import (
	"gorm.io/gorm"
)

func newConn() *gorm.DB{
	g, err := conn(); 
	if err == nil {
		return g
	}
	panic(err)
}