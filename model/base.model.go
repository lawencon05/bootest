package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
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

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	return nil
}
