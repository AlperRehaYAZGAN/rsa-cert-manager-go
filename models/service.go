package model

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null;unique"`
	Key  string `gorm:"type:varchar(64);not null;unique"`
}

// CreateServiceDto is the data transfer object for creating a service.
type CreateServiceDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// UpdateServiceDto is the data transfer object for updating a service.
type UpdateServiceDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}

// DeleteServiceDto is the data transfer object for deleting a service.
type DeleteServiceDto struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
	Key  string `json:"key" binding:"required,min=1,max=64"`
}
