package main

import (
	"info7255-bigdata-app/routes"
	"log"
)

func main() {
	r := routes.SetupRouter()

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
