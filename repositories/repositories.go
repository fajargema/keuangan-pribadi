package repositories

import "keuangan-pribadi/models"

type UserRepository interface {
	Register(UserInput models.UserInput) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Login(UserInput models.UserAuth) (models.UserResponse, error)
	Update(UserInput models.UserInput, token string) (models.User, error)
}