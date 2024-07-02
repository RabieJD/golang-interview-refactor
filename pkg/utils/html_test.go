package utils_test

import (
	"github.com/stretchr/testify/assert"
	td "interview/pkg/testdata/utils"
	"interview/pkg/utils"
	"os"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	for _, tt := range td.GenerateTTRenderTemplate() {
		t.Run(tt.Name, func(t *testing.T) {
			// Create template file if it doesn't exist
			if tt.Source == td.SourceHtmlTestFile {
				err := os.WriteFile(tt.Source, []byte("<h1>{{.Title}}</h1>"), 0644)
				assert.NoError(t, err)
			}
			t.Cleanup(func() {
				if tt.Source == td.SourceHtmlTestFile {
					err := os.Remove(tt.Source)
					assert.NoError(t, err)
				}
			})
			result, err := utils.RenderTemplate(tt.PageData, tt.Source)

			if tt.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, result)
			}
		})
	}
}
