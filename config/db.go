package config

import (
	"fmt"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lawencon.com/bootest/model"
)

var tables = []interface{}{
	&model.AnswerDtl{},
	&model.AnswerHdr{},
	&model.CandidateProfiles{},
	&model.EmployeeProfiles{},
	&model.QuestionTypes{},
	&model.Questions{},
	&model.Roles{},
	&model.Users{},
}

const (
	host    = "localhost"
	user    = "postgres"
	pass    = "postgres"
	dbname  = "bootest"
	port    = 5432
	sslmode = "disable"
)

func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, pass, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
