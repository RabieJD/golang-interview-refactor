package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetRequestForm(c *gin.Context, form interface{}) error {
	if c.Request.Body == nil {
		return fmt.Errorf("body cannot be nil")
	}
	return binding.FormPost.Bind(c.Request, form)
}
