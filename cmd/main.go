package main

import (
	"github.com/dotenv-org/godotenvvault"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"os"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	dotenvErr := godotenvvault.Load()
	if dotenvErr != nil {
		logger.Fatal("failed to load .env file", zap.Error(dotenvErr))
	}

	gormLogger := zapgorm2.New(logger)
	gormLogger.SetAsDefault()

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dsn := "host=" + host + " user=" + user + " pass=" + pass + " dbname=" + name + " port=" + port

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if dbErr != nil {
		logger.Fatal("failed to connect to database", zap.Error(dbErr))
	}
}
