package models

import "time"

type BaseModels struct {
	IsActive    bool
	CreatedBy   string
	UpdatedBy   string
	CreatedDate time.Time
	UpdatedDate time.Time
}
