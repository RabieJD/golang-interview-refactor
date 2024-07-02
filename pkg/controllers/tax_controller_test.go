package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interview/pkg/controllers"
	td "interview/pkg/testdata/controllers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTaxController_AddItem(t *testing.T) {
	for _, tt := range td.GenerateTTAddItem() {
		t.Run(tt.Name, func(t *testing.T) {
			mockHandler := &td.CalculatorHandlerMock{}
			tt.Setup(mockHandler)

			taxController := controllers.NewTaxController(mockHandler)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			if tt.Cookie != nil {
				req.AddCookie(tt.Cookie)
			}
			c.Request = req

			c.Request.PostForm = url.Values{}
			for k, v := range tt.FormData {
				c.Request.PostForm.Set(k, v)
			}

			taxController.AddItem(c)
			assert.Equal(t, tt.Expected, c.Writer.Status())
		})
	}
}

func TestTaxController_ShowAddItemForm(t *testing.T) {
	for _, tt := range td.GenerateTTTShowAddItemForm() {
		t.Run(tt.Name, func(t *testing.T) {
			mockHandler := &td.CalculatorHandlerMock{}
			tt.Setup(mockHandler)

			taxController := controllers.NewTaxController(mockHandler)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.Cookie != nil {
				req.AddCookie(tt.Cookie)
			}
			c.Request = req

			if len(tt.QueryParams) > 0 {
				q := c.Request.URL.Query()
				for k, v := range tt.QueryParams {
					q.Add(k, v)
				}
				c.Request.URL.RawQuery = q.Encode()
			}

			taxController.ShowAddItemForm(c)
			assert.Equal(t, tt.Expected, c.Writer.Status())
		})
	}
}

func TestTaxController_DeleteCartItem(t *testing.T) {
	for _, tt := range td.GenerateTTDeleteItem() {
		t.Run(tt.Name, func(t *testing.T) {
			mockHandler := &td.CalculatorHandlerMock{}
			tt.Setup(mockHandler)

			taxController := controllers.NewTaxController(mockHandler)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			if tt.Cookie != nil {
				req.AddCookie(tt.Cookie)
			}
			c.Request = req

			q := c.Request.URL.Query()
			for k, v := range tt.QueryParams {
				q.Add(k, v)
			}
			c.Request.URL.RawQuery = q.Encode()

			taxController.DeleteCartItem(c)
			assert.Equal(t, tt.Expected, c.Writer.Status())
		})
	}
}
