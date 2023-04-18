package repositories

import (
	"keuangan-pribadi/config"
	m "keuangan-pribadi/middleware"
	"keuangan-pribadi/models"
	"time"
)

type FinanceRepositoryImpl struct{}

func InitFinanceRepository() FinanceRepository {
	return &FinanceRepositoryImpl{}
}

func (fr *FinanceRepositoryImpl) GetAll(token string) ([]models.Finance, error) {
	var finances []models.Finance

	user, err := m.VerifyToken(token)
    if err != nil {
        return []models.Finance{}, err
    }

	if err := config.DB.Where("user_id = ?", user.ID).Preload("User").Preload("Category").Find(&finances).Error; err != nil {
		return nil, err
	}

	return finances, nil
}

func (fr *FinanceRepositoryImpl) GetByID(id, token string) (models.Finance, error) {
	var finance models.Finance

	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Finance{}, err
    }

	if err := config.DB.Preload("User").Preload("Category").First(&finance, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return models.Finance{}, err
	}

	return finance, nil
}

func (fr *FinanceRepositoryImpl) Search(from, to time.Time, token string) ([]models.Finance, error) {
	var finances []models.Finance

	user, err := m.VerifyToken(token)
    if err != nil {
        return []models.Finance{}, err
    }

	if err := config.DB.Where("created_at BETWEEN ? AND ? AND user_id = ?", from, to, user.ID).Preload("User").Preload("Category").Find(&finances).Error; err != nil {
		return nil, err
	}

	return finances, nil
}

func (fr *FinanceRepositoryImpl) Create(financeInput models.FinanceInput, token string) (models.Finance, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Finance{}, err
    }

	var User models.User
	if err := config.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
		return models.Finance{}, err
	}

	var category models.Category
	if err := config.DB.Where("id = ?", financeInput.CategoryID).First(&category).Error; err != nil {
		return models.Finance{}, err
	}

	var createdFinance models.Finance = models.Finance{
		Name:       	financeInput.Name,
		Type: 			financeInput.Type,
		Money: 			financeInput.Money,
		UserID:    		user.ID,
		CategoryID:    	financeInput.CategoryID,
		User: 			User,
		Category: 		category,
	}

	result := config.DB.Create(&createdFinance)

	if err := result.Error; err != nil {
		return models.Finance{}, err
	}

	if err := config.DB.Last(&createdFinance).Error; err != nil {
		return models.Finance{}, err
	}

	return createdFinance, nil
}

func (fr *FinanceRepositoryImpl) Update(financeInput models.FinanceInput, id, token string) (models.Finance, error) {
	user, err := m.VerifyToken(token)
    if err != nil {
        return models.Finance{}, err
    }

	finance, err := fr.GetByID(id, token)
	if err != nil {
		return models.Finance{}, err
	}

	var User models.User
	if err := config.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
		return models.Finance{}, err
	}

	var category models.Category
	if err := config.DB.Where("id = ?", financeInput.CategoryID).First(&category).Error; err != nil {
		return models.Finance{}, err
	}

	finance.Name = financeInput.Name
	finance.Type = financeInput.Type
	finance.Money = financeInput.Money
	finance.UserID = user.ID
	finance.CategoryID = financeInput.CategoryID
	finance.User = User
	finance.Category = category

	if err := config.DB.Save(&finance).Error; err != nil {
		return models.Finance{}, err
	}

	return finance, nil
}

func (fr *FinanceRepositoryImpl) Delete(id, token string) error {
	_, err := m.VerifyToken(token)
    if err != nil {
        return err
    }

	finance, err := fr.GetByID(id, token)

	if err != nil {
		return err
	}

	if err := config.DB.Delete(&finance).Error; err != nil {
		return err
	}

	return nil
}