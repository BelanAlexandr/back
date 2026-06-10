package main

import (
	"log"

	"github.com/BelanAlexandr/back/internal/config"
	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/BelanAlexandr/back/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	repository.InitDB(cfg)
	router := routes.Routes()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
