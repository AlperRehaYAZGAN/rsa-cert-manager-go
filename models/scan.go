package model

import (
	"gorm.io/gorm"
)

type Scan struct {
	gorm.Model
	Key string `gorm:"type:varchar(100);unique_index"`
}
