package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
        "time"

	"test-app/internal/database"
	"test-app/internal/handlers"
        "test-app/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	// 🔥 ДОБАВЛЕНО: Prometheus metrics
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gofiber/adaptor/v2"
)

// 🔥 ДОБАВЛЕНО: Кастомные метрики
var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func main() {
	godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	// 🔥 ДОБАВЛЕНО: Graceful shutdown для БД
	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("Database connection closed")
		}
	}()

	db.AutoMigrate(&models.Task{})

	app := fiber.New(fiber.Config{
		AppName: "DevOps App v1.0",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	
	// 🔥 ДОБАВЛЕНО: Middleware для сбора метрик
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()
		
		status := c.Response().StatusCode()
		httpRequestsTotal.WithLabelValues(c.Method(), c.Path(), string(rune(status))).Inc()
		httpRequestDuration.WithLabelValues(c.Method(), c.Path()).Observe(duration)
		
		return err
	})

	// 🔥 ИСПРАВЛЕНО: Настоящий Prometheus endpoint
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	handlers.RegisterHealthHandlers(app)
	handlers.RegisterTaskHandlers(app, db)

	// 🔥 ДОБАВЛЕНО: Корневой endpoint (редирект на /health или API docs)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "DevOps App",
			"version": "1.0",
			"endpoints": fiber.Map{
				"health":    "/health",
				"ready":     "/ready",
				"metrics":   "/metrics",
				"api_tasks": "/api/tasks",
			},
		})
	})

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

	address := host + ":" + port
	log.Printf("Server starting on %s", address)
	if err := app.Listen(address); err != nil {
		log.Fatal("failed to listen:", err)
	}
}
