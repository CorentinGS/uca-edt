package controllers

import (
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func GetStudentEDT(c *fiber.Ctx) error {
	uuid := c.Params("id")
	answer := database.GetEdt(uuid)

	return c.JSON(answer)

}
