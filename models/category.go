package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        	uint           	`json:"id" gorm:"primaryKey"`
	Name     	string 		 	`json:"name" form:"name"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type CategoryInput struct {
	Name     	string 	`json:"name" form:"name" validate:"required"`
}