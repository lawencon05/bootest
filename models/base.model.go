package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

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

func (base *BaseModels) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	base.Id = uuid.String()
	return nil
}
