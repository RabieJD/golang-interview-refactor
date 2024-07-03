package config_test

import (
	"github.com/stretchr/testify/assert"
	"interview/pkg/config"
	td "interview/pkg/testdata/config"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := td.GenerateTTNew()
	// save already existing env
	keys := []string{
		"SERVER_HOST", "SERVER_PORT", "MYSQL_DATABASE", "MYSQL_HOST", "MYSQL_PORT", "MYSQL_USER", "MYSQL_PASSWORD",
	}
	oldEnv := map[string]string{}
	for _, k := range keys {
		oldEnv[k] = os.Getenv(k)
		_ = os.Unsetenv(k)
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			// set env
			for k, v := range tt.Env {
				oldEnv[k] = os.Getenv(k)
				_ = os.Setenv(k, v)
			}
			// create config
			conf := config.New()
			// check config
			assert.EqualValues(t, tt.ExpectedResult, conf)
		})
	}
	// revert old env
	for k, v := range oldEnv {
		_ = os.Setenv(k, v)
	}
}
