package config

import (
	"log"
)

type Config struct {
	MongoURI string
	MongoDB  string
}

func LoadConfig() Config {
	log.Println("Loading configuration...")
	cfg := Config{
		MongoURI: "mongodb://localhost:27017",
		MongoDB:  "event_sourcing_db",
	}
	return cfg
}
