package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormServicePostgreImpl struct{
	*gorm.DB
}

const (
	host    = "localhost"
	user    = "postgres"
	pass    = "postgres"
	dbname  = "bootest"
	port    = 5432
	sslmode = "disable"
)

func (gormService *GormServicePostgreImpl) Conn() error {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta", host, user, pass, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if(err != nil){
		return err
	}
	
	*gormService.DB = *db
	return nil
}
