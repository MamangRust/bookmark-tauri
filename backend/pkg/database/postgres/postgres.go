package postgres

import (
	"bookmark-backend/internal/models"
	"bookmark-backend/pkg/logger"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(log logger.Logger) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
	// 	viper.GetString("DB_USER"),
	// 	viper.GetString("DB_PASSWORD"),
	// 	viper.GetString("DB_HOST"),
	// 	viper.GetString("DB_NAME"),
	// )

	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_USER:", user)
	fmt.Println("DB_PASSWORD:", password)
	fmt.Println("DB_NAME:", dbname)
	fmt.Println("DB_PORT:", port)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", viper.GetString("DB_HOST"), viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"), viper.GetString("DB_PORT"))

	// dsn := "host=localhost user=holyraven password=holyraven dbname=bookmark_db port=5435 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Migrate User and Post tables first
	if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Error("Failed to auto migrate User and Post models", zap.Error(err))
		defer log.Info("Database Connection failed")
		return nil, fmt.Errorf("failed to auto migrate User and Post models: %v", err)
	}

	// Migrate Category table after User and Post tables exist
	if err := db.AutoMigrate(&models.Category{}); err != nil {
		log.Error("Failed to auto migrate Category model", zap.Error(err))
		defer log.Info("Database Connection failed")
		return nil, fmt.Errorf("failed to auto migrate Category model: %v", err)
	}

	return db, nil
}
