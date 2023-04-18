package config

import (
	"fmt"
	"keuangan-pribadi/utils"
	"log"

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
	//
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