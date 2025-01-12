package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ConfigDatabase
}

type ConfigDatabase struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func InitConfig(configPath string) (*Config, error) {
	err := godotenv.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err)
	}

	host := os.Getenv("AUTH_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	portStr := os.Getenv("AUTH_DB_PORT")
	if portStr == "" {
		portStr = "5432"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port in config: %s", err)
	}

	username := os.Getenv("AUTH_DB_USERNAME")
	passwd := os.Getenv("AUTH_DB_PASSWORD")
	dbname := os.Getenv("AUTH_DB_DBNAME")

	return &Config{
		ConfigDatabase {
			Host: host,
			Port: port,
			Username: username,
			Password: passwd,
			DatabaseName: dbname,
		},
	}, nil
}