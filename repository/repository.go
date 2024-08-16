package repository

import (
	"fmt"

	"bbb/models"

	"os"

	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DbName   string
	SslMode  string
}

func (config DBConfig) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DbName, config.SslMode)
}

func Create(candidate string) (uint, error) {
	db, err := openDB()

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	list := models.Vote{Candidate: candidate}

	result := db.Create(&list)

	if result.Error != nil {
		return 0, result.Error
	}

	return list.ID, nil
}

func Migrate() {
	db, err := openDB()
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Vote{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func openDB() (*gorm.DB, error) {
	godotenv.Load()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	config := DBConfig{
		User:     "postgres",
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     port,
		DbName:   os.Getenv("DATABASE"),
		SslMode:  "disable",
	}

	db, err := gorm.Open(postgres.Open(config.ConnectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetConnMaxLifetime(300)
	return db, nil
}
