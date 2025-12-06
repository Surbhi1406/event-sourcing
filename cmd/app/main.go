package main

import (
	cfg "eventSourcing/config"
	"eventSourcing/internal/database"
	"eventSourcing/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := cfg.LoadConfig()
	dbClient, db, err := database.ConnectToDB(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)

	}
	defer dbClient.Disconnect(nil)

	r := gin.New()

	// Setup all routes with database dependency
	routes.SetupRoutes(r, db)

	r.Run()
}
