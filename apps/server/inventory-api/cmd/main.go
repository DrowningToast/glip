package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	go func() {
		log.Println("Server is starting...")
		err := app.Listen(":3000")
		if err != nil {
			log.Fatal(err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("Server process is interrupted")

	go func() {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		go func() {
			<-timeoutCtx.Done()
			if timeoutCtx.Err() == context.DeadlineExceeded {
				log.Println("Server shutdown with timeout")
				os.Exit(1)
			}
		}()

		if err := app.ShutdownWithContext(timeoutCtx); err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			log.Println("Server shutdown gracefully")
		}
	}()
}
