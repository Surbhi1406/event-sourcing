package routes

import (
	"eventSourcing/internal/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	// Initialize handlers with dependencies
	healthHandler := handlers.NewHealthHandler()
	orderHandler := handlers.NewOrderHandler(db)

	// Health check route
	router.GET("/", healthHandler.HealthCheck)

	// Order routes
	router.POST("/createOrder", orderHandler.CreateOrder)
	router.GET("/getOrders", orderHandler.GetOrders)
	router.POST("/updateOrder", orderHandler.UpdateOrder)
	router.DELETE("/deleteOrder", orderHandler.DeleteOrder)
}
