package services

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func InitUserService() UserService {
	return UserService{
		repository: &repositories.UserRepositoryImpl{},
	}
}

func (us *UserService) GetByEmail(email string) (models.User, error) {
	return us.repository.GetByEmail(email)
}

func (us *UserService) Register(userInput models.UserInput) (models.User, error) {
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput models.UserAuth) (models.UserResponse, error) {
	return us.repository.Login(userInput)
}

func (us *UserService) UpdateMe(userInput models.UserInput, token string) (models.User, error) {
	return us.repository.UpdateMe(userInput, token)
}