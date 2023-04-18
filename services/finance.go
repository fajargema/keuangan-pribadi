package services

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/repositories"
	"time"
)

type FinanceService struct {
	repository repositories.FinanceRepository
}

func InitFinanceService() FinanceService {
	return FinanceService{
		repository: &repositories.FinanceRepositoryImpl{},
	}
}

func (fs *FinanceService) GetAll(token string) ([]models.Finance, error) {
	return fs.repository.GetAll(token)
}

func (fs *FinanceService) GetByID(id, token string) (models.Finance, error) {
	return fs.repository.GetByID(id, token)
}

func (fs *FinanceService) Search(from, to time.Time, token string) ([]models.Finance, error) {
	return fs.repository.Search(from, to, token)
}

func (fs *FinanceService) Create(financeInput models.FinanceInput, token string) (models.Finance, error) {
	return fs.repository.Create(financeInput, token)
}

func (fs *FinanceService) Update(financeInput models.FinanceInput, id, token string) (models.Finance, error) {
	return fs.repository.Update(financeInput, id, token)
}

func (fs *FinanceService) Delete(id, token string) error {
	return fs.repository.Delete(id, token)
}
