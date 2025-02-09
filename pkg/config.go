package pkg

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type SYSConfig struct {
	SYS_PORT        string
	DB_USER         string
	DB_PASSWORD     string
	DB_HOST         string
	DB_PORT         string
	DB_NAME         string
	DB_SSLMODE      string
	JWT_SECRET      string
	UPLOAD_MAX_SIZE string
	UPLOAD_DIR      string
	REDIS_HOST      string
	REDIS_PASSWORD  string
	REDIS_DB        string
	OSS_ENDPOINT    string
	OSS_ACCESS_KEY  string
	OSS_SECRET_KEY  string
	OSS_BUCKET      string
	OSS_PREFIX_URL  string
}

func LoadConfig() SYSConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return SYSConfig{
		SYS_PORT:        GetEnv("SYS_PORT", ""),
		DB_USER:         GetEnv("DB_USER", ""),
		DB_PASSWORD:     GetEnv("DB_PASSWORD", ""),
		DB_HOST:         GetEnv("DB_HOST", ""),
		DB_PORT:         GetEnv("DB_PORT", ""),
		DB_NAME:         GetEnv("DB_NAME", ""),
		DB_SSLMODE:      GetEnv("DB_SSLMODE", ""),
		JWT_SECRET:      GetEnv("JWT_SECRET", ""),
		UPLOAD_MAX_SIZE: GetEnv("UPLOAD_MAX_SIZE", ""),
		UPLOAD_DIR:      GetEnv("UPLOAD_DIR", ""),
		REDIS_HOST:      GetEnv("REDIS_HOST", ""),
		REDIS_PASSWORD:  GetEnv("REDIS_PASSWORD", ""),
		REDIS_DB:        GetEnv("REDIS_DB", ""),
		OSS_ENDPOINT:    GetEnv("OSS_ENDPOINT", ""),
		OSS_ACCESS_KEY:  GetEnv("OSS_ACCESS_KEY", ""),
		OSS_SECRET_KEY:  GetEnv("OSS_SECRET_KEY", ""),
		OSS_BUCKET:      GetEnv("OSS_BUCKET", ""),
		OSS_PREFIX_URL:  GetEnv("OSS_PREFIX_URL", ""),
	}
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
