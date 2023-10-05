package controllers

import (
	"d0c/TestTaskBackDev/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "d0c/TestTaskBackDev/docs"

	"github.com/gofiber/swagger"
)

func NewRouter(app *fiber.App, services *services.Services) {
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(200) })
	app.Get("/swagger/*", swagger.HandlerDefault)

	auth := app.Group("/auth")
	{
		newAuthRoutes(auth, services.Session)
	}
}
