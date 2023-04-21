package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailSaving struct {
	ID        	uint           	`json:"id" gorm:"primaryKey"`
	Value     	int 		 	`json:"value" form:"value"`
	SavingID 	uint 			`json:"saving_id" form:"saving_id"`
	Saving   	Saving 			`gorm:"foreignKey:SavingID"`
	UserID 		uint 			`json:"user_id" form:"user_id"`
	User   		User 			`gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type DetailSavingInput struct {
	Value     	int 	`json:"value" form:"value" validate:"required"`
	SavingID    uint 	`json:"saving_id" form:"saving_id" validate:"required"`
	UserID 		uint 	`json:"user_id" form:"user_id"`
}