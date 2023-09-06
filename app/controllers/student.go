package controllers

import (
	"net/http"

	"github.com/corentings/uca-edt/pkg/core"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/corentings/uca-edt/pkg/models"
	"github.com/corentings/uca-edt/pkg/parsing"
	"github.com/gofiber/fiber/v2"
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

// PostStudentEDT posts a student edt from a given UUID
func PostStudentEDT(c *fiber.Ctx) error {
	// uuid param
	uuid := c.Params("id")

	// studentCourses body
	studentCourses := new(map[string]string)

	// Parse body
	if err := c.BodyParser(studentCourses); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while parsing body: " + err.Error(),
		})
	}

	// Get courseEdt from database
	courseEdt, err := database.GetCourseData()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Course edt not found: " + err.Error(),
		})
	}

	// Create a student
	student := make(map[string][]models.CourseEDT, 1)

	// Compute student edt
	studentEDT := core.ComputeStudent(models.StudentJSON{Courses: *studentCourses}, (*parsing.CourseEdt)(&courseEdt))
	student[uuid] = studentEDT

	// Store student edt in database
	database.StoreEdt(student)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student " + uuid + " edt posted",
	})
}
