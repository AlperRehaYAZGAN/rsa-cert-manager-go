package model

import (
	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null;unique"`
	Key  string `gorm:"type:varchar(64);not null;unique"`
}

// CreatePlatformDto is the data transfer object for creating a platform.
type CreatePlatformDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// UpdatePlatformDto is the data transfer object for updating a platform.
type UpdatePlatformDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// DeletePlatformDto is the data transfer object for deleting a platform.
type DeletePlatformDto struct {
	Key string `json:"key" binding:"required,min=1,max=64"`
}
