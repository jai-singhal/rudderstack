package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./.env/.local")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		serverPort = 8080
	}

	serverMode := os.Getenv("SERVER_MODE")
	if serverMode == "" {
		serverMode = "debug"
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 3306
	}

	config := &Config{
		Server: ServerConfig{
			Port: serverPort,
			Mode: serverMode,
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	return config, nil
}
