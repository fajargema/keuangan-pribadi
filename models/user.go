package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        	uint           	`json:"id" gorm:"primaryKey"`
	Name     	string 			`json:"name" form:"name"`
	Email    	string 			`json:"email" form:"email"`
	Password 	string 			`json:"password" form:"password"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type UserInput struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type UserAuth struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type UserResponse struct {
	ID 		uint 	`json:"id" form:"id"`
	Name    string 	`json:"name" form:"name"`
	Email   string 	`json:"email" form:"email"`
	Token 	string 	`json:"token" form:"token"`
}