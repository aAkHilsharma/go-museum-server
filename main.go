package main

import (
	"log"
	"museum/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/exhibitions", api.Get)
	router.POST("/api/exhibitions", api.Post)

	router.Static("/api/assets", "./public")

	err := router.Run(":3333")
	if err != nil {
		log.Fatal(err)
	}
}
