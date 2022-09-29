package api

import (
	"fmt"
	"os"

	models "github.com/TheaKevin/helloworld/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "l1a2p3t4o5p6"
	dbname   = "tutorialGoReact"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		config := fmt.Sprintf("host=%s port=%d user =%s password =%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := db.AutoMigrate(&models.Todos{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, err
}
