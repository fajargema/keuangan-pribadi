package models

import (
	"time"

	"gorm.io/gorm"
)

type Saving struct {
	ID        	uint           	`json:"id" gorm:"primaryKey"`
	Name     	string 		 	`json:"name" form:"name"`
	Value     	int 		 	`json:"value" form:"value"`
	Goal     	int 		 	`json:"goal" form:"goal"`
	UserID 		uint 			`json:"user_id" form:"user_id"`
	User   		User 			`gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type SavingInput struct {
	Name     	string 	`json:"name" form:"name" validate:"required"`
	Value     	int 	`json:"value" form:"value"`
	Goal     	int 	`json:"goal" form:"goal" validate:"required"`
}

type SavingUpdate struct {
	Name     	string 	`json:"name" form:"name" validate:"required"`
	Goal     	int 	`json:"goal" form:"goal" validate:"required"`
}