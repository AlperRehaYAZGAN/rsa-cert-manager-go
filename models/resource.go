package model

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Key string `gorm:"type:varchar(100);unique_index"`
}
