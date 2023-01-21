package api

import (
	"github.com/gin-gonic/gin"
	"hex-architecture-boilerplate/infrastructure/adapter/api/handler"
)

func Router() {
	router := gin.Default()
	router.GET("/", handler.HealthyCheck)
	router.Run()
}
