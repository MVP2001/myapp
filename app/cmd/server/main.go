package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"test-app/internal/database"
	"test-app/internal/handlers"
	"test-app/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	db.AutoMigrate(&models.Task{})

	app := fiber.New(fiber.Config{
		AppName: "DevOps App v1.0",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	handlers.RegisterHealthHandlers(app)
	handlers.RegisterTaskHandlers(app, db)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		app.Shutdown()
	}()

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// 🔥 ИСПРАВЛЕНО: правильная конкатенация с двоеточием
	address := host + ":" + port
	log.Printf("Server starting on %s", address)
	if err := app.Listen(address); err != nil {
		log.Fatal("failed to listen:", err)
	}
}
