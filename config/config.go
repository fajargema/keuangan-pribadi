package config

import (
	"fmt"
	"keuangan-pribadi/models"
	"keuangan-pribadi/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	InitDB()
}

var (
	DB *gorm.DB
	DB_USERNAME string = utils.GetConfig("DB_USERNAME")
	DB_PASSWORD string = utils.GetConfig("DB_PASSWORD")
	DB_NAME     string = utils.GetConfig("DB_NAME")
	DB_HOST     string = utils.GetConfig("DB_HOST")
	DB_PORT     string = utils.GetConfig("DB_PORT")
)

// connect to the database
func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Finance{}, &models.Saving{}, &models.DetailSaving{})
}

func SeedUser() (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte("testsecret"), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name: "test",
		Email: "test@gmail.com",
		Password: string(password),
	}

	result := DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SeedCategory() (models.Category, error) {
	var category models.Category = models.Category{
		Name: "seederform",
	}

	result := DB.Create(&category)

	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	if err := result.Last(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func SeedFinance() (models.Finance, error) {
	user, err := SeedUser()
	if err != nil {
		return models.Finance{}, err
	}

	category, err := SeedCategory()
	if err != nil {
		return models.Finance{}, err
	}

	var finance models.Finance = models.Finance{
		Name:       	"test",
		Type: 			1,
		Money: 			10000,
		UserID:  		user.ID,
		CategoryID:  	category.ID,
		User:       	user,
		Category:       category,
	}

	result := DB.Create(&finance)

	if err := result.Error; err != nil {
		return models.Finance{}, err
	}

	if err := result.Last(&finance).Error; err != nil {
		return models.Finance{}, err
	}

	return finance, nil
}

func SeedSaving() (models.Saving, error) {
	user, err := SeedUser()
	if err != nil {
		return models.Saving{}, err
	}

	var saving models.Saving = models.Saving{
		Name:       	"test",
		Value: 			1,
		Goal: 			10000,
		UserID:  		user.ID,
		User:       	user,
	}

	result := DB.Create(&saving)

	if err := result.Error; err != nil {
		return models.Saving{}, err
	}

	if err := result.Last(&saving).Error; err != nil {
		return models.Saving{}, err
	}

	return saving, nil
}

func SeedDetailSaving() (models.DetailSaving, error) {
	user, err := SeedUser()
	if err != nil {
		return models.DetailSaving{}, err
	}

	saving, err := SeedSaving()
	if err != nil {
		return models.DetailSaving{}, err
	}

	var detailSaving models.DetailSaving = models.DetailSaving{
		Value: 			1,
		SavingID: 			saving.ID,
		UserID:  		user.ID,
		User:       	user,
		Saving:       	saving,
	}

	result := DB.Create(&detailSaving)

	if err := result.Error; err != nil {
		return models.DetailSaving{}, err
	}

	if err := result.Last(&detailSaving).Error; err != nil {
		return models.DetailSaving{}, err
	}

	return detailSaving, nil
}

func CloseDB() error {
	database, err := DB.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}