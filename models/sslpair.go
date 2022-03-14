package model

import (
	"gorm.io/gorm"
)

type SSLPair struct {
	gorm.Model
	Key        string `gorm:"type:varchar(100);not null;unique"`
	PublicKey  string `gorm:"type:varchar(4096)"`
	PrivateKey string `gorm:"type:varchar(4096)"`
}

// CreateSSLPairDto is the data transfer object for creating a ssl pair.
type CreateSSLPairDto struct {
	Key        string `json:"key" binding:"required,min=1,max=100"`
	PublicKey  string `json:"public_key" binding:"required,min=1,max=4096"`
	PrivateKey string `json:"private_key" binding:"required,min=1,max=4096"`
}

// UpdateSSLPairDto is the data transfer object for updating a ssl pair.
type UpdateSSLPairDto struct {
	Key        string `json:"key" binding:"required,min=1,max=100"`
	PublicKey  string `json:"public_key" binding:"required,min=1,max=4096"`
	PrivateKey string `json:"private_key" binding:"required,min=1,max=4096"`
}

// DeleteSSLPairDto is the data transfer object for deleting a ssl pair.
type DeleteSSLPairDto struct {
	Key string `json:"key" binding:"required,min=1,max=100"`
}
