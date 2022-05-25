package main

import (
	"log"
	config "todo_app/internal/configs"
	"todo_app/internal/di"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatalf("calling Load: %v", err)
	}

	server := di.InitializeServer(config)

	err = server.Run()
	if err != nil {
		log.Fatalf("calling Run: %v", err)
	}
}
