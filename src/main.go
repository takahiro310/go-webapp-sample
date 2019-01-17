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
	router.GET("/db-select", controller.DBGet)
	router.Run(":8080")
}
