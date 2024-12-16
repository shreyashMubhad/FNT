package main

import (
	"DC/FnO/cmd/app"
	"DC/FnO/pkg/config"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	var environment string

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <env>")
		os.Exit(1)
	}

	env := os.Args[1]
	if env != "uat" {
		environment = "prod"
	} else {
		environment = "uat"
	}
	fmt.Println("ENV: ", environment)
	config.Load(environment)

	ctx := context.Background()

	// Start the server
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}