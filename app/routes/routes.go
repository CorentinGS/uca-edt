package routes

import (
	"github.com/bytedance/sonic"
	"github.com/corentings/uca-edt/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func New() *fiber.App {
	// Create new app
	app := fiber.New(
		fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, *",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 2
	}))

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   10 * time.Minute,
		CacheControl: true,
	}))

	// Api group
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	api.Get("/edt/:id", controllers.GetStudentEDT)

	return app
}
