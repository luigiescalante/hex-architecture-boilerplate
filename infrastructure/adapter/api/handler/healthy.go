package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthyCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"system": "hexagonal architecture in GO",
		"date":   time.Now(),
	})
}
