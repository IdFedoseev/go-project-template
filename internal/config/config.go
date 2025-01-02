package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Server     ServerConfig
    MongoDB    MongoDBConfig
    Environment string
    Telemetry   TelemetryConfig
}

type ServerConfig struct {
    Port string
}

type MongoDBConfig struct {
    URI      string
    Database string
}

type TelemetryConfig struct {
    Host string
    Port string
}

func Load() *Config {
    // Загружаем .env файл
    if err := godotenv.Load(); err != nil {
        // Если файл не найден, используем переменные окружения
        // Это полезно для production окружения, где .env может не использоваться
    }

    return &Config{
        Server: ServerConfig{
            Port: getEnv("SERVER_PORT", "8080"),
        },
        MongoDB: MongoDBConfig{
            URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
            Database: getEnv("MONGO_DATABASE", "keldi"),
        },
        Environment: getEnv("ENVIRONMENT", "development"),
        Telemetry: TelemetryConfig{
            Host: getEnv("OTEL_COLLECTOR_HOST", "localhost"),
            Port: getEnv("OTEL_COLLECTOR_PORT", "4318"),
        },
    }
}

// getEnv получает значение из переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
} 