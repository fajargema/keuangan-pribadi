package repositories

import (
	"keuangan-pribadi/config"
	m "keuangan-pribadi/middleware"
	"keuangan-pribadi/models"
)

type SavingRepositoryImpl struct{}

func InitSavingRepository() SavingRepository {
	return &SavingRepositoryImpl{}
}

func (sr *SavingRepositoryImpl) GetAll(token string) ([]models.Saving, error) {
	var savings []models.Saving

	user, err := m.VerifyToken(token)
    if err != nil {
        return []models.Saving{}, err
    }

	if err := config.DB.Preload("User").Where("user_id = ?", user.ID).Find(&savings).Error; err != nil {
		return nil, err
	}

	return savings, nil
}

func (sr *SavingRepositoryImpl) GetByID(id, token string) (models.Saving, error) {
	var saving models.Saving

	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Saving{}, err
    }

	if err := config.DB.Preload("User").First(&saving, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return models.Saving{}, err
	}

	return saving, nil
}

func (sr *SavingRepositoryImpl) Create(savingInput models.SavingInput, token string) (models.Saving, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Saving{}, err
    }

	var User models.User
	er := config.DB.Where("id = ?", user.ID).First(&User).Error
	if er != nil {
		return models.Saving{}, er
	}

	var createdSaving models.Saving = models.Saving{
		Name:       	savingInput.Name,
		Value: 			savingInput.Value,
		Goal: 			savingInput.Goal,
		UserID:    		user.ID,
		User: 			User,
	}

	result := config.DB.Create(&createdSaving)

	if err := result.Error; err != nil {
		return models.Saving{}, err
	}

	if err := config.DB.Last(&createdSaving).Error; err != nil {
		return models.Saving{}, err
	}

	var createdDetailSaving models.DetailSaving = models.DetailSaving{
		Value: 			savingInput.Value,
		Status: 		1,
		UserID:    		user.ID,
		User: 			User,
		SavingID: 		createdSaving.ID,
		Saving: 		createdSaving,
	}

	if err := config.DB.Create(&createdDetailSaving).Error; err != nil{
		return models.Saving{}, err
	}

	if err := config.DB.Last(&createdDetailSaving).Error; err != nil {
		return models.Saving{}, err
	}

	return createdSaving, nil
}

func (sr *SavingRepositoryImpl) Update(savingUpdate models.SavingUpdate, id, token string) (models.Saving, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Saving{}, err
    }

	saving, err := sr.GetByID(id, token)
	if err != nil {
		return models.Saving{}, err
	}

	var User models.User
	er := config.DB.Where("id = ?", user.ID).First(&User).Error
	if er != nil {
		return models.Saving{}, er
	}

	saving.Name = savingUpdate.Name
	saving.Goal = savingUpdate.Goal
	saving.UserID = user.ID
	saving.User = User

	if err := config.DB.Save(&saving).Error; err != nil {
		return models.Saving{}, err
	}

	return saving, nil
}

func (sr *SavingRepositoryImpl) Delete(id, token string) error {
	_, err := m.VerifyToken(token)
    if err != nil {
        return err
    }

	saving, err := sr.GetByID(id, token)

	if err != nil {
		return err
	}

	if err := config.DB.Delete(&saving).Error; err != nil {
		return err
	}

	return nil
}