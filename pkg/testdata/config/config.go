package config

import (
	"interview/pkg/config"
	"interview/pkg/infra/database/connection"
)

// TTNew is a test case for a NEwConfig function
type TTNew struct {
	Name           string
	Env            map[string]string
	ExpectedResult *config.Config
}

// GenerateTTNew generates test cases for the New config function
func GenerateTTNew() []TTNew {
	return []TTNew{
		{
			Name: "default config",
			Env:  make(map[string]string),
			ExpectedResult: &config.Config{
				Server: config.Server{
					Host: "localhost",
					Port: "8088",
				},
				DatabaseConnection: connection.DatabaseConnection{
					Name:     "ice_db",
					Host:     "localhost",
					Port:     "4001",
					Username: "user",
					Password: "user",
				},
			},
		},
		{
			Name: "config from env",
			Env: map[string]string{
				"SERVER_HOST":    "test-host",
				"SERVER_PORT":    "1234",
				"MYSQL_DATABASE": "test-db-name",
				"MYSQL_HOST":     "test-host",
				"MYSQL_PORT":     "5678",
				"MYSQL_USER":     "test-user",
				"MYSQL_PASSWORD": "test-password",
			},
			ExpectedResult: &config.Config{
				Server: config.Server{
					Host: "test-host",
					Port: "1234",
				},
				DatabaseConnection: connection.DatabaseConnection{
					Name:     "test-db-name",
					Host:     "test-host",
					Port:     "5678",
					Username: "test-user",
					Password: "test-password",
				},
			},
		},
	}
}
