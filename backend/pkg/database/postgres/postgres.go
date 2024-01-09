package postgres

import (
	"bookmark-backend/internal/models"
	"bookmark-backend/pkg/logger"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(log logger.Logger) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", viper.GetString("DB_HOST"), viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"), viper.GetString("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Error("Failed to get database connection", zap.Error(err))
		defer log.Info("Database Connection failed")
		return nil, fmt.Errorf("failed to get database connection: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}); err != nil {
		log.Error("Failed to auto migrate User, Category and Post models", zap.Error(err))

		return nil, fmt.Errorf("failed to auto migrate User, Category and Post models: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
