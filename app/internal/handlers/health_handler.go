package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterHealthHandlers(app *fiber.App) {
	// Liveness probe (проверка, что сервер жив)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "test-app",
		})
	})

	// Readiness probe (проверка готовности принимать трафик)
	app.Get("/ready", func(c *fiber.Ctx) error {
		// Здесь можно добавить проверку подключения к БД
		return c.JSON(fiber.Map{
			"status": "ready",
		})
	})
}
