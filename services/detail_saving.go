package services

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/repositories"
)

type DetailSavingService struct {
	repository repositories.DetailSavingRepository
}

func InitDetailSavingService() DetailSavingService {
	return DetailSavingService{
		repository: &repositories.DetailSavingRepositoryImpl{},
	}
}

func (dss *DetailSavingService) GetAll(token string) ([]models.DetailSaving, error) {
	return dss.repository.GetAll(token)
}

func (dss *DetailSavingService) GetByID(id, token string) (models.DetailSaving, error) {
	return dss.repository.GetByID(id, token)
}

func (dss *DetailSavingService) Create(detailSavingInput models.DetailSavingInput, token string) (models.DetailSaving, error) {
	return dss.repository.Create(detailSavingInput, token)
}

func (dss *DetailSavingService) Update(detailSavingInput models.DetailSavingInput, id, token string) (models.DetailSaving, error) {
	return dss.repository.Update(detailSavingInput, id, token)
}

func (dss *DetailSavingService) Delete(id, token string) error {
	return dss.repository.Delete(id, token)
}
