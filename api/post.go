package api

import (
	"museum/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var exhibition data.Exhibition
	err := c.ShouldBindJSON(&exhibition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	data.AddExhibition(exhibition)
	c.JSON(http.StatusCreated, exhibition)
}
