package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ConfigDatabase
	ConfigJWT
}

type ConfigDatabase struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

type ConfigJWT struct {
	SecretKey string
	ExpAt int
}

func MustInitConfig(configPath string) *Config {
	err := godotenv.Load(configPath)
	if err != nil {
		panic("error loading .env file")
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
		panic("invalid port in config")
	}

	username := os.Getenv("AUTH_DB_USERNAME")
	passwd := os.Getenv("AUTH_DB_PASSWORD")
	dbname := os.Getenv("AUTH_DB_DBNAME")

	secretKey := os.Getenv("AUTH_SECRET_KEY")

	expAtStr := os.Getenv("AUTH_EXP_AT")
	if expAtStr == "" {
		expAtStr = "1"
	}

	expAt, _ := strconv.Atoi(expAtStr)

	return &Config{
		ConfigDatabase {
			Host: host,
			Port: port,
			Username: username,
			Password: passwd,
			DatabaseName: dbname,
		},
		ConfigJWT{
			SecretKey: secretKey,
			ExpAt: expAt,	
		},
	}
}