package api

import (
	"museum/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Query("id")

	if id != "" {
		finalId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		c.JSON(http.StatusOK, data.GetAll()[finalId])
		return
	}

	c.JSON(http.StatusOK, data.GetAll())
}
