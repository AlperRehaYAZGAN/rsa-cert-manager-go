package model

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null;unique"`
	Key  string `gorm:"type:varchar(64);not null;unique"`
}

// CreateResourceDto is the data transfer object for creating a resource.
type CreateResourceDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// UpdateResourceDto is the data transfer object for updating a resource.
type UpdateResourceDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// DeleteResourceDto is the data transfer object for deleting a resource.
type DeleteResourceDto struct {
	Key string `json:"key" binding:"required,min=1,max=64"`
}
