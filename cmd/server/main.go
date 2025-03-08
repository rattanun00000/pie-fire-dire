package main

import (
	"log"

	"pie-fire-dire/internal/app"
)

func main() {
	log.Println("Initializing application...")
	application := app.NewApp()

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err := application.Run(port)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
