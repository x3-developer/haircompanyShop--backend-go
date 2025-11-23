package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppEnv        string
	AppPort       string
	DbHost        string
	DbPort        string
	DbName        string
	DbUser        string
	DbPassword    string
	DbSsl         string
	CORS          string
	AuthAppKey    string
	AppLogLvl     string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	AppSecret     string
}

func LoadConfig() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST environment variable is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("DB_PORT environment variable is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME environment variable is not set")

	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER environment variable is not set")

	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is not set")

	}

	dbSsl := os.Getenv("DB_SSL")
	if dbSsl == "" {
		dbSsl = "verify-full"
	}

	corsAllowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if corsAllowedOrigins == "" {
		return nil, fmt.Errorf("CORS_ALLOWED_ORIGINS environment isn't set")

	}

	authAppKey := os.Getenv("AUTH_APP_KEY")
	if authAppKey == "" {
		return nil, fmt.Errorf("AUTH_APP_KEY environment isn't set")
	}

	appLogLvl := os.Getenv("APP_LOG_LEVEL")
	if appLogLvl == "" {
		appEnv = "debug"
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		return nil, fmt.Errorf("REDIS_ADDR environment variable is not set")
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		return nil, fmt.Errorf("REDIS_DB environment variable is not set")
	}
	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Fatal("Invalid REDIS_DB value: ", err)
	}

	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		return nil, fmt.Errorf("APP_SECRET environment variable is not set")
	}

	return &Config{
		AppEnv:        appEnv,
		AppPort:       appPort,
		DbHost:        dbHost,
		DbPort:        dbPort,
		DbName:        dbName,
		DbUser:        dbUser,
		DbPassword:    dbPassword,
		DbSsl:         dbSsl,
		CORS:          corsAllowedOrigins,
		AuthAppKey:    authAppKey,
		AppLogLvl:     appLogLvl,
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
		RedisDB:       redisDBInt,
		AppSecret:     appSecret,
	}, nil
}
