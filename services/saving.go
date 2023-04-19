package services

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/repositories"
)

type SavingService struct {
	repository repositories.SavingRepository
}

func InitSavingService() SavingService {
	return SavingService{
		repository: &repositories.SavingRepositoryImpl{},
	}
}

func (ss *SavingService) GetAll(token string) ([]models.Saving, error) {
	return ss.repository.GetAll(token)
}

func (ss *SavingService) GetByID(id, token string) (models.Saving, error) {
	return ss.repository.GetByID(id, token)
}

func (ss *SavingService) Create(savingInput models.SavingInput, token string) (models.Saving, error) {
	return ss.repository.Create(savingInput, token)
}

func (ss *SavingService) Update(savingUpdate models.SavingUpdate, id, token string) (models.Saving, error) {
	return ss.repository.Update(savingUpdate, id, token)
}

func (ss *SavingService) Delete(id, token string) error {
	return ss.repository.Delete(id, token)
}
