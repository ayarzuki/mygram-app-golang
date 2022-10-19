package dbconfig

import (
	"fmt"
	"log"
	"mygram-app-golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(driver string) (db *gorm.DB, err error) {
	if driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", GoDotEnvVariable("DB_HOST"), GoDotEnvVariable("DB_PORT"), GoDotEnvVariable("DB_USER"), GoDotEnvVariable("DB_PASS"), GoDotEnvVariable("DB_NAME"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("driver/dbms %s is not supported", driver)
	}

	log.Default().Println("DB connected successfully")
	db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
	return db, nil
}
