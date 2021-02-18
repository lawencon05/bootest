package config

import (
	"gorm.io/gorm"

	"lawencon.com/bootest/db"
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

var gormService db.GormService = db.GormServicePostgreImpl{}

func Conn() (*gorm.DB, error) {
	db, err := gormService.Conn()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
