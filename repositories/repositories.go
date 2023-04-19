package repositories

import (
	"keuangan-pribadi/models"
	"time"
)

type UserRepository interface {
	Register(UserInput models.UserInput) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Login(UserInput models.UserAuth) (models.UserResponse, error)
	UpdateMe(UserInput models.UserInput, token string) (models.User, error)
}

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(CategoryInput models.CategoryInput) (models.Category, error)
	Update(CategoryInput models.CategoryInput, id string) (models.Category, error)
	Delete(id string) error
}

type FinanceRepository interface {
	GetAll(token string) ([]models.Finance, error)
	GetByID(id, token string) (models.Finance, error)
	Search(from, to time.Time, token string) ([]models.Finance, error)
	Create(FinanceInput models.FinanceInput, token string) (models.Finance, error)
	Update(FinanceInput models.FinanceInput, id, token string) (models.Finance, error)
	Delete(id, token string) error
}

type SavingRepository interface {
	GetAll(token string) ([]models.Saving, error)
	GetByID(id, token string) (models.Saving, error)
	Create(SavingInput models.SavingInput, token string) (models.Saving, error)
	Update(SavingUpdate models.SavingUpdate, id, token string) (models.Saving, error)
	Delete(id, token string) error
}