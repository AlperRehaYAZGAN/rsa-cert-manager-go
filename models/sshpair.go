package model

import (
	"gorm.io/gorm"
)

type SSHPair struct {
	gorm.Model
	Key        string `gorm:"type:varchar(100);not null;unique"`
	PublicKey  string `gorm:"type:varchar(4096)"`
	PrivateKey string `gorm:"type:varchar(4096)"`
}

// CreateSSHPairDto is the data transfer object for creating a ssh pair.
type CreateSSHPairDto struct {
	Key        string `json:"key" binding:"required,min=1,max=100"`
	PublicKey  string `json:"public_key" binding:"required,min=1,max=4096"`
	PrivateKey string `json:"private_key" binding:"required,min=1,max=4096"`
}

// UpdateSSHPairDto is the data transfer object for updating a ssh pair.
type UpdateSSHPairDto struct {
	Key        string `json:"key" binding:"required,min=1,max=100"`
	PublicKey  string `json:"public_key" binding:"required,min=1,max=4096"`
	PrivateKey string `json:"private_key" binding:"required,min=1,max=4096"`
}

// DeleteSSHPairDto is the data transfer object for deleting a ssh pair.
type DeleteSSHPairDto struct {
	Key string `json:"key" binding:"required,min=1,max=100"`
}
