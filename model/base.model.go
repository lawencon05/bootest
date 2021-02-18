package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"lawencon.com/bootest/util"
)

type BaseModel struct {
	Id          string  `json:"id" gorm:"primaryKey;type:varchar(50)"`
	IsActive    bool    `gorm:"default:true"`
	CreatedBy   string  `json:"createdBy"`
	UpdatedBy   *string `json:"updatedBy"`
	CreatedDate string  `json:"createdDate" gorm:"type:timestamp"`
	UpdatedDate *string `json:"updatedDate" gorm:"type:timestamp"`
	Count       int64   `gorm:"-"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	base.CreatedDate = util.TimeToStr(time.Now())
	return nil
}
