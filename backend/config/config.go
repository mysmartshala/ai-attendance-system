package config

import (
	"os"
)

type Config struct {
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	ServerPort          string
	JWTSecret           string
	FaceServiceURL      string
	MaxFileSize         int64
	UploadPath          string
	SimilarityThreshold float32
}

func LoadConfig() *Config {
	return &Config{
		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnv("DB_PORT", "3306"),
		DBUser:              getEnv("DB_USER", "root"),
		DBPassword:          getEnv("DB_PASSWORD", "root"),
		DBName:              getEnv("DB_NAME", "attendance"),
		ServerPort:          getEnv("SERVER_PORT", "8080"),
		JWTSecret:           getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		FaceServiceURL:      getEnv("FACE_SERVICE_URL", "http://localhost:5000"),
		MaxFileSize:         50 * 1024 * 1024,
		UploadPath:          getEnv("UPLOAD_PATH", "./uploads"),
		SimilarityThreshold: 0.60,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
