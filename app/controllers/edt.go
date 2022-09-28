package controllers

import (
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetCourseData(c *fiber.Ctx) error {

	// Get student edt from database
	answer, err := database.GetCourseData()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Course edt not found: " + err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(answer)
}
