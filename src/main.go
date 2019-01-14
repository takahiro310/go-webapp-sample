package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.IndexGet)

	router.GET("/json-get", controller.JSONGet)
	router.POST("/json-post", controller.JSONPost)

	router.Run(":8080")
}
