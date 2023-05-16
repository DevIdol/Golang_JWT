package config

import (
	"fmt"
	"time"

	"github.com/DevIdol/Golang_JWT/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.ErrorPanic(err)

	sqlDB, err := db.DB()
	helpers.ErrorPanic(err)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("PostgresDB Connected!")

	return db, nil
}