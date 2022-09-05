package core

import (
	"github.com/corentings/uca-edt/pkg/models"
	"github.com/corentings/uca-edt/pkg/parsing"
)

func ComputeStudentEDT() {
	student := models.StudentJSON{
		UUID: "101748",
		Courses: map[string]string{
			"MMAG":   "2",
			"MF":     "1",
			"Chimie": "1",
		},
	}

	edt := parsing.GetCourseEdt()

	for key, value := range student.Courses {
		for key2, course := range edt[key].CourseEDT {
			if course.Groupe == value || course.Type == "CM" {
				edt[key].CourseEDT[key2].Print()
			}
		}
	}
}
