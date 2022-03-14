package model

import (
	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	Key  string `gorm:"type:varchar(100);unique_index"`
}
