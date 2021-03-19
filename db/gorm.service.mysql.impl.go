package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormServiceMysqlImpl struct{
	*gorm.DB
}

const (
	hostMysql   = "localhost"
	userMysql   = "root"
	passMysql   = "mysql"
	dbnameMysql = "bootest"
	portMysql   = 3306
	charset     = "utf8"
)

func (gormService *GormServiceMysqlImpl) Conn() error {
	pg := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		userMysql, passMysql, hostMysql, portMysql, dbnameMysql, charset)

	db, err := gorm.Open(mysql.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if(err != nil){
		return err
	}
	
	gormService.DB = db
	return nil
}
