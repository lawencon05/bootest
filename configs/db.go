package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lawencon.com/imamfarisi/models"
)

var tables = []interface{}{
	&models.AnswerDtl{},
	&models.AnswerHdr{},
	&models.CandidateProfiles{},
	&models.EmployeeProfiles{},
	&models.QuestionTypes{},
	&models.Questions{},
	&models.Roles{},
	&models.Users{},
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
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
