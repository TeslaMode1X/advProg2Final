package db

import "os"

type Database struct {
	User         string `env:"DB_USER, required"`
	Password     string `env:"DB_PASSWORD, required"`
	Host         string `env:"DB_HOST, required"`
	Port         string `env:"DB_PORT, required"`
	DriverName   string `env:"DB_DRIVER, required"`
	DatabaseName string `env:"DB_NAME, required"`
	SSLMode      string `env:"DB_SSLMODE, default=disable"`
}

func InitDbConfig() Database {
	return Database{
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_NAME"),
		DriverName:   os.Getenv("DB_DRIVER"),
		SSLMode:      os.Getenv("DB_SSLMODE"),
	}
}
