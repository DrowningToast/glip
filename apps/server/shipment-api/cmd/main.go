package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create background context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Shipment API",
	})

	// Basic route
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Shipment API!")
	})

	// Start server with graceful shutdown
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	// Cancel context
	cancel()
}
