package handlers

import (
	"test-app/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterTaskHandlers(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/tasks")

	// CREATE
	api.Post("/", func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		
		if err := db.Create(task).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		
		return c.Status(201).JSON(task)
	})

	// READ ALL
	api.Get("/", func(c *fiber.Ctx) error {
		var tasks []models.Task
		if err := db.Find(&tasks).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(tasks)
	})

	// READ ONE
	api.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var task models.Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.JSON(task)
	})

	// UPDATE
	api.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var task models.Task
		
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
		}
		
		updates := new(models.Task)
		if err := c.BodyParser(updates); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		
		db.Model(&task).Updates(updates)
		return c.JSON(task)
	})

	// DELETE
	api.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&models.Task{}, id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	})
}
