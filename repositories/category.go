package repositories

import (
	"keuangan-pribadi/config"
	"keuangan-pribadi/models"
)

type CategoryRepositoryImpl struct{}

func InitCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (cr *CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := config.DB.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepositoryImpl) GetByID(id string) (models.Category, error) {
	var category models.Category

	err := config.DB.First(&category, "id = ?", id).Error

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Create(categoryInput models.CategoryInput) (models.Category, error) {
	var createdCategory models.Category = models.Category{
		Name:       categoryInput.Name,
	}

	result := config.DB.Create(&createdCategory)

	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	err := config.DB.Last(&createdCategory).Error

	if err != nil {
		return models.Category{}, err
	}

	return createdCategory, nil
}

func (cr *CategoryRepositoryImpl) Update(categoryInput models.CategoryInput, id string) (models.Category, error) {
	category, err := cr.GetByID(id)

	if err != nil {
		return models.Category{}, err
	}

	category.Name = categoryInput.Name

	err = config.DB.Save(&category).Error

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Delete(id string) error {
	category, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	err = config.DB.Delete(&category).Error

	if err != nil {
		return err
	}

	return nil
}