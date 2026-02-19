package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"test-app/internal/database"
	"test-app/internal/handlers"

	"github.com/gofiber/contrib/fiberprometheus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env только в development
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
	}

	// Подключение к БД
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Auto-migration
	if err := db.AutoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Получаем sql.DB для graceful shutdown
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
	}

	// Настройка Fiber с таймаутами
	app := fiber.New(fiber.Config{
		AppName:      "DevOps App v1.0",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Prometheus метрики
	prometheus := fiberprometheus.New("myapp")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		// Проверяем соединение с БД
		if err := sqlDB.Ping(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "unhealthy",
				"error":  "database connection failed",
			})
		}
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
		})
	})

	// Регистрация обработчиков
	handlers.RegisterHealthHandlers(app)
	handlers.RegisterTaskHandlers(app, db)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")

		// Контекст с таймаутом для shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Закрываем соединение с БД
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}

		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	// Определяем host и port
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	address := host + ":" + port
	log.Printf("Server starting on %s", address)

	if err := app.Listen(address); err != nil {
		log.Fatal("Failed to listen:", err)
	}
}
