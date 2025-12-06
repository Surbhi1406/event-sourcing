package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

// NewHealthHandler creates a new HealthHandler instance
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck handles health check endpoint
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	log.Println("Handling msg")
	c.JSON(http.StatusOK, gin.H{"msg": "Hello Surbhi"})
}
