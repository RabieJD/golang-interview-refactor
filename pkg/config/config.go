package config

import (
	"interview/pkg/infra/database/connection"
	"os"
)

const (
	// default server configuration
	defaultServerPort = "8088"
	defaultServerHost = "localhost"
	// default database connection configuration
	defaultDatabaseName = "ice_db"
	defaultDatabaseUser = "user"
	defaultDatabasePass = "user"
	defaultDatabaseHost = "localhost"
	defaultDatabasePort = "4001"
)

// Config is a container for all the needed app configuration.
type Config struct {
	Server             Server
	DatabaseConnection connection.DatabaseConnection
}

// Server holds the server configuration.
type Server struct {
	Host string `default:"localhost" env:"SERVER_HOST"`
	Port string `default:"8088" env:"SERVER_PORT"`
}

// Storage holds the storage configuration.
type Storage struct {
	Type string `default:"memory" env:"STORAGE_TYPE"`
}

// New initialize the config.
func New() *Config {
	return &Config{
		Server: Server{
			Host: getOrDefault("SERVER_HOST", defaultServerHost),
			Port: getOrDefault("SERVER_PORT", defaultServerPort),
		},
		DatabaseConnection: connection.DatabaseConnection{
			Name:     getOrDefault("MYSQL_DATABASE", defaultDatabaseName),
			Host:     getOrDefault("MYSQL_HOST", defaultDatabaseHost),
			Port:     getOrDefault("MYSQL_PORT", defaultDatabasePort),
			Username: getOrDefault("MYSQL_USER", defaultDatabaseUser),
			Password: getOrDefault("MYSQL_PASSWORD", defaultDatabasePass),
		},
	}
}
func getOrDefault(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
