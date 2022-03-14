package model

import (
	"gorm.io/gorm"
)

type Scan struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null;unique"`
	Key  string `gorm:"type:varchar(64);not null;unique"`
	// value is jsonb postgresql type
	Value string `gorm:"type:jsonb"`
}
