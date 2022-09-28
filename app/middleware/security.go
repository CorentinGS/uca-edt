package middleware

import "github.com/gofiber/fiber/v2"

var SecurityKey string

// SecurityMiddleware is the middleware for the security

func IsSecurityKeyValid(key string) bool {
	return key == SecurityKey
}

func SecurityKeyMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if !IsSecurityKeyValid(c.Get("Key")) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid security key",
			})
		}
		return c.Next()
	}
}
