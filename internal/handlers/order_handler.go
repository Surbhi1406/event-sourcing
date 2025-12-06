package handlers

import (
	"log"
	"net/http"

	"eventSourcing/internal/database"
	"eventSourcing/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewOrder struct {
	OrderID  string
	Item     string
	Quantity int
	Price    float64
}
type OrderHandler struct {
	db *mongo.Database
}

// NewOrderHandler creates a new OrderHandler instance with database dependency
func NewOrderHandler(db *mongo.Database) *OrderHandler {
	ord := OrderHandler{
		db: db,
	}
	return &ord
}

// CreateOrder handles order creation
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	log.Println("Handling createOrderHandler")

	// Now you can access the database via h.db
	// Example: collection := h.db.Collection("orders")

	var newOrder models.Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalud Json"})
		return
	}

	log.Printf("Recived Order Payload %v\n", newOrder)
	err := database.AddOrderInDB(c.Request.Context(), h.db, newOrder)
	if err != nil {
		log.Println("Failed to add order to DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Order added successfully"})
}

// GetOrders handles fetching orders
func (h *OrderHandler) GetOrders(c *gin.Context) {
	log.Println("Handling getOrderHandler")
	orders, err := database.GetOrdersFromDB(c.Request.Context(), h.db)
	if err != nil {
		log.Println("Failed to get orders from DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// UpdateOrder handles order updates
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	log.Println("Handling updateOrderHandler")
	// TODO: Implement update order logic
}

// DeleteOrder handles order deletion
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	log.Println("Handling deleteOrderHandler")
	// TODO: Implement delete order logic
}
