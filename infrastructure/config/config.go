package config

import (
    "os"
    "strconv"
    "strings"
)

// Config represent config keys.
type Config struct {
    App
    DBConfig
}

// New returns a new Config struct.
func New() *Config {
    return &Config{
        App: App{
            Name: getEnv("APP_NAME", "GO Rest API"),
            Port: getEnv("APP_PORT", "8080"),
        },
        DBConfig: DBConfig{
            DBDriver:   getEnv("DB_DRIVER", "mysql"),
            DBHost:     getEnv("DB_HOST", "localhost"),
            DBPort:     getEnv("DB_POST", "3306"),
            DBUser:     getEnv("DB_USER", "root"),
            DBName:     getEnv("DB_NAME", "go_rest_skeleton"),
            DBPassword: getEnv("DB_PASSWORD", ""),
            DBLog:      getEnvAsBool("ENABLE_LOGGER", true),
        },
    }
}

// Simple helper function to read an environment or return a default value.
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    if nextValue := os.Getenv(key); nextValue != "" {
        return nextValue
    }

    return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }

    return defaultVal
}

// Helper to read an environment variable into a bool or return default value.
func getEnvAsBool(name string, defaultVal bool) bool {
    valStr := getEnv(name, "")
    if val, err := strconv.ParseBool(valStr); err == nil {
        return val
    }

    return defaultVal
}

// Helper to read an environment variable into a string slice or return default value.
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
    valStr := getEnv(name, "")

    if valStr == "" {
        return defaultVal
    }

    val := strings.Split(valStr, sep)

    return val
}
