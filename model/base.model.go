package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Timestamp time.Time

type BaseModel struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	IsActive    bool       `gorm:"default:true"`
	CreatedBy   string     `json:"createdBy"`
	UpdatedBy   *string    `json:"updatedBy"`
	CreatedDate *Timestamp `json:"createdDate" gorm:"type:timestamp without time zone;"`
	UpdatedDate *Timestamp `json:"updatedDate" gorm:"type:timestamp without time zone;"`
	Count       int64      `gorm:"-"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	return nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return nil
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	layout := "2006-01-02 15:04:05"

	ts, err := time.Parse(layout, timeStr)
	*t = Timestamp(ts)

	return err
}
