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
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			// set env
			for k, v := range tt.Env {
				_ = os.Setenv(k, v)
			}
			// create config
			conf := config.New()
			// check config
			assert.EqualValues(t, conf, tt.ExpectedResult)
		})
	}
}
