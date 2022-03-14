package model

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Key string `gorm:"type:varchar(100);unique_index"`
}
