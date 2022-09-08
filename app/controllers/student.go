package controllers

import (
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetStudentEDT returns a student edt from a given UUID
func GetStudentEDT(c *fiber.Ctx) error {
	// uuid param
	uuid := c.Params("id")

	// Get student edt from database
	answer, err := database.GetEdt(uuid)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Student " + uuid + " not found: " + err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(answer)
}
