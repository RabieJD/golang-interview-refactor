package utils_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	td "interview/pkg/testdata/utils"
	"interview/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequestForm(t *testing.T) {
	for _, tt := range td.GenerateTTGetRequestForm() {
		t.Run(tt.Name, func(t *testing.T) {
			// Set up the Gin context with the test body
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = &http.Request{
				Method: "POST",
				Header: http.Header{
					"Content-Type": []string{"application/x-www-form-urlencoded"},
				},
				Body: tt.Body,
			}

			var form td.TTForm
			err := utils.GetRequestForm(c, &form)

			if tt.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, form)
			}
		})
	}
}
