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

func SetDb() *gorm.DB {
	var g = &gorm.DB{}
	var gormService db.GormService = &db.GormServicePostgreImpl{
		DB: g,
	}
	err := gormService.Conn()

	if err != nil {
		return nil
	}

	g.AutoMigrate(tables...)
	return g
}
