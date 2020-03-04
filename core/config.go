package core

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var AppConfig *Config

// init is invoke before main()
func init() {
	// load values from .env into the system
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "./..")
	if err := godotenv.Load(basePath + "/.env"); err != nil {
		log.Println("No .env file found")
	}
	AppConfig = NewConfig()
}

type Config struct {
	Debug                bool
	Testing              bool
	DisableNotifications bool
	DatabaseURL          string
	SECRET               string
	AllowedHosts         []string
	Port                 string
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	return &Config{
		Debug:                getEnvAsBool("DEBUG", true),
		Testing:              getEnvAsBool("TESTING", false),
		DisableNotifications: getEnvAsBool("DISABLE_NOTIFICATIONS", true),
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		SECRET:               getEnv("SECRET_KEY", ""),
		AllowedHosts:         getEnvAsSlice("ALLOWED_HOSTS", []string{"*"}, ","),
		Port:                 getEnv("PORT", "8080"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into an integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
