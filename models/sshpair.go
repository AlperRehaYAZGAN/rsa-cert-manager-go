package model

import (
	"gorm.io/gorm"
)

type SSHPair struct {
	gorm.Model
	Key        string `gorm:"type:varchar(100);unique_index"`
	PublicKey  string `gorm:"type:varchar(4096)"`
	PrivateKey string `gorm:"type:varchar(4096)"`
}
