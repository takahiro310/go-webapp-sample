package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Msg is JSON struct
type Msg struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

// Error is request error result
type Error struct {
	Number  int
	Message string
}

// JSONGet is json data I/O sample
func JSONGet(c *gin.Context) {

	var msg Msg
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123

	c.JSON(http.StatusOK, msg)
}

// JSONPost is json data I/O sample
func JSONPost(c *gin.Context) {

	var msg Msg
	if c.BindJSON(&msg) == nil {
		c.JSON(http.StatusOK, msg)
	}
	// Bad Request
	var error Error
	error.Number = 400
	error.Message = "パラメータが不正です"
	c.JSON(http.StatusExpectationFailed, error)
}
