package configs

import (
	"assignment/helpers"
	"assignment/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := helpers.GetEnv("DB_HOST")
	user := helpers.GetEnv("DB_USER")
	pass := helpers.GetEnv("DB_PASS")
	dbname := helpers.GetEnv("DB_DBNAME")
	port := helpers.GetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, pass, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.PanicIfError(err, "DB connection problem")

	db.AutoMigrate(&models.Order{}, &models.Item{})

	DB = db
}
