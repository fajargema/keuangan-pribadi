package repositories

import (
	"fmt"
	"keuangan-pribadi/config"
	m "keuangan-pribadi/middleware"
	"keuangan-pribadi/models"
)

type DetailSavingRepositoryImpl struct{}

func InitDetailSavingRepository() DetailSavingRepository {
	return &DetailSavingRepositoryImpl{}
}

func (dsr *DetailSavingRepositoryImpl) GetAll(token string) ([]models.DetailSaving, error) {
	var detailSavings []models.DetailSaving

	user, err := m.VerifyToken(token)
    if err != nil {
        return []models.DetailSaving{}, err
    }

	if err := config.DB.Preload("User").Preload("Saving.User").Where("user_id = ?", user.ID).Find(&detailSavings).Error; err != nil {
		return nil, err
	}

	return detailSavings, nil
}

func (dsr *DetailSavingRepositoryImpl) GetByID(id, token string) (models.DetailSaving, error) {
	var detailSaving models.DetailSaving

	user, err := m.VerifyToken(token)
    if err != nil {
        return models.DetailSaving{}, err
    }

	if err := config.DB.Preload("User").Preload("Saving.User").First(&detailSaving, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return models.DetailSaving{}, err
	}

	return detailSaving, nil
}

func (dsr *DetailSavingRepositoryImpl) Create(savingInput models.DetailSavingInput, token string) (models.DetailSaving, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.DetailSaving{}, err
    }

	var User models.User
	if err := config.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
		return models.DetailSaving{}, err
	}

	var Saving models.Saving
	if err := config.DB.Preload("User").Where("id = ?", savingInput.SavingID).First(&Saving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	total := Saving.Value + savingInput.Value
	if err := config.DB.Model(&Saving).Update("value", total).Error; err != nil {
		return models.DetailSaving{}, err
	}

	if total >= Saving.Goal {
		exp := 10 + User.Exp
		if User.Exp == exp {
			exp := exp
			fmt.Println("ini exp:", exp)
			if err := config.DB.Model(&User).Update("exp", exp).Error; err != nil {
				return models.DetailSaving{}, err
			}
		} else if User.Exp < exp {
			exp := User.Exp
			fmt.Println("ini user.exp:", exp)
			if err := config.DB.Model(&User).Update("exp", exp).Error; err != nil {
				return models.DetailSaving{}, err
			}
		}
	}
	
	if err := config.DB.Preload("User").Where("id = ?", savingInput.SavingID).First(&Saving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	var createdDetailSaving models.DetailSaving = models.DetailSaving{
		Value: 			savingInput.Value,
		UserID:    		user.ID,
		SavingID:    	savingInput.SavingID,
		User: 			User,
		Saving: 		Saving,
	}

	result := config.DB.Create(&createdDetailSaving)

	if err := result.Error; err != nil {
		return models.DetailSaving{}, err
	}

	if err := config.DB.Last(&createdDetailSaving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	return createdDetailSaving, nil
}

func (dsr *DetailSavingRepositoryImpl) Update(savingInput models.DetailSavingInput, id, token string) (models.DetailSaving, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.DetailSaving{}, err
    }

	detailSaving, err := dsr.GetByID(id, token)
	if err != nil {
		return models.DetailSaving{}, err
	}

	var User models.User
	if err := config.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
		return models.DetailSaving{}, err
	}

	var Saving models.Saving
	if err := config.DB.Preload("User").Where("id = ?", savingInput.SavingID).First(&Saving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	kurang := Saving.Value - detailSaving.Value
	if err := config.DB.Model(&Saving).Update("value", kurang).Error; err != nil {
		return models.DetailSaving{}, err
	}

	total := Saving.Value + savingInput.Value
	if err := config.DB.Model(&Saving).Update("value", total).Error; err != nil {
		return models.DetailSaving{}, err
	}

	detailSaving.Value = savingInput.Value
	detailSaving.UserID = user.ID
	detailSaving.SavingID = savingInput.SavingID
	detailSaving.User = User
	detailSaving.Saving = Saving

	if err := config.DB.Save(&detailSaving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	return detailSaving, nil
}

func (dsr *DetailSavingRepositoryImpl) Delete(id, token string) error {
	_, err := m.VerifyToken(token)
    if err != nil {
        return err
    }

	detailSaving, err := dsr.GetByID(id, token)

	var Saving models.Saving
	if err := config.DB.Preload("User").Where("id = ?", detailSaving.SavingID).First(&Saving).Error; err != nil {
		return err
	}

	kurang := Saving.Value - detailSaving.Value
	if err := config.DB.Model(&Saving).Update("value", kurang).Error; err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if err := config.DB.Delete(&detailSaving).Error; err != nil {
		return err
	}

	return nil
}