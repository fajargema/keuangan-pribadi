package models

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	ID        	uint           	`json:"id" gorm:"primaryKey"`
	Name     	string 			`json:"name" form:"name"`
	Type		int 			`json:"type" form:"type" gorm:"check:type IN(1,2)"`
	Money		int 			`json:"money" form:"money"`
	UserID 		uint 			`json:"user_id" form:"user_id"`
	CategoryID 	uint 			`json:"category_id" form:"category_id"`
	User   		User 			`gorm:"foreignKey:UserID"`
	Category   	Category 		`gorm:"foreignKey:CategoryID"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type FinanceInput struct {
	Name		string 	`json:"name" form:"name" validate:"required"`
	Type    	int 	`json:"type" form:"type" validate:"required"`
	Money    	int 	`json:"money" form:"money" validate:"required"`
	UserID 		uint 	`json:"user_id" form:"user_id"`
	CategoryID 	uint 	`json:"category_id" form:"category_id" validate:"required"`
}