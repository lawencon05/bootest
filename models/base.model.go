package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BaseModels struct {
	Id          string  `json:"id" gorm:"primaryKey"`
	IsActive    bool    `gorm:"default:true"`
	CreatedBy   string  `json:"createdBy"`
	UpdatedBy   *string `json:"updatedBy"`
	CreatedDate time.Time
	UpdatedDate *time.Time
	Count       int64 `gorm:"-"`
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
