package config

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/statistics/config/db"
	"github.com/TeslaMode1X/advProg2Final/statistics/config/server"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

const ReviewEnvName = "statistics"

type Config struct {
	DB     *db.Database
	Server *server.Server
}

func InitConfig() Config {
	err := loadDotEnv()
	if err != nil {
		log.Printf("Warning: failed to load .env file: %v. Proceeding with defaults or env vars.", err)
	}

	srv := server.InitServerConfig()
	db := db.InitDbConfig()

	return Config{
		Server: &srv,
		DB:     &db,
	}
}

func loadDotEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Println("Current working directory:", dir)

	filePath := fmt.Sprintf(".env.%s", ReviewEnvName)

	if _, err := os.Stat(filePath); err == nil {
		return godotenv.Load(filePath)
	}

	filePath = filepath.Join("../..", fmt.Sprintf(".env.%s", ReviewEnvName))
	return godotenv.Load(filePath)
}
