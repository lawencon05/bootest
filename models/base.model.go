package models

import "time"

type BaseModels struct {
	Id       string `json:"id" gorm:"primary_key"`
	IsActive bool
	// CreatedBy   string
	// UpdatedBy   string
	CreatedDate time.Time
	// UpdatedDate time.Time
	Count int64 `sql:"-" gorm:"-"`
}

type Tabler interface {
	TableName() string
}
