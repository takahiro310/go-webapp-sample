package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGet is root document
func IndexGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
