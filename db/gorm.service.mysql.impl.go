package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormServiceMysqlImpl struct{}

const (
	hostMysql   = "localhost"
	userMysql   = "root"
	passMysql   = "mysql"
	dbnameMysql = "bootest"
	portMysql   = 3306
	charset     = "utf8"
)

func (GormServiceMysqlImpl) Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		userMysql, passMysql, hostMysql, portMysql, dbnameMysql, charset)

	return gorm.Open(mysql.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
