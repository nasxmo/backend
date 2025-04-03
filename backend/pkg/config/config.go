package config

import (
	"os"
)

type Config struct {
	DBURL           string
	KratosURL       string
	KratosAdminURL  string
	HydraURL        string
	HydraAdminURL   string
	OathkeeperURL   string
	Port            string
	SessionSecret   string
	CookieDomain    string
}

func Load() *Config {
	return &Config{
		DBURL:           getEnv("DB_URL", "postgres://ory:secret@postgres:5432/ory?sslmode=disable"),
		KratosURL:       getEnv("ORY_KRATOS_URL", "http://kratos:4433"),
		KratosAdminURL:  getEnv("ORY_KRATOS_ADMIN_URL", "http://kratos:4434"),
		HydraURL:        getEnv("ORY_HYDRA_URL", "http://hydra:4444"),
		HydraAdminURL:   getEnv("ORY_HYDRA_ADMIN_URL", "http://hydra:4445"),
		OathkeeperURL:   getEnv("ORY_OATHKEEPER_URL", "http://oathkeeper:4455"),
		Port:            getEnv("PORT", "8080"),
		SessionSecret:   getEnv("SESSION_SECRET", "your-secret-key-here"),
		CookieDomain:    getEnv("COOKIE_DOMAIN", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}