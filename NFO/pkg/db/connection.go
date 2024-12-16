package db

import (
	"DC/FnO/pkg/config"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Renaming the struct to PostgresDB
type PostgresDB struct {
	logger *zap.Logger
	DB     *gorm.DB
}

// Function to initialize a new PostgresDB
func NewPostgresDB(cfg *config.Config, logger *zap.Logger) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode)

	logger.Info("connecting to database")
	// Open the database connection using gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to the database", zap.Error(err))
		return nil, err
	}

	// Return the initialized PostgresDB struct
	return &PostgresDB{DB: db, logger: logger}, nil
}
