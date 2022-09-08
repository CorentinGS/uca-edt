package routes

import (
	"github.com/bytedance/sonic"
	"github.com/corentings/uca-edt/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() *fiber.App {
	// Create new app
	app := fiber.New(
		fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		})

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, *",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	// Compression
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 2
	}))

	// Api group
	api := app.Group("/api")

	// Routes
	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	api.Get("/edt/:id", controllers.GetStudentEDT) // Get student edt

	return app
}
