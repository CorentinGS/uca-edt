package parsing

import "github.com/corentings/uca-edt/pkg/models"

type CourseEdt map[string]models.CourseData
type DataEdtJSON map[string]map[string][]models.Course
type DataStudentJSON map[string]map[string]string

func GetCourseEdt() CourseEdt {
	data := parseEdtJSON("parsed.json")
	edt := parseEdt(*data)
	return *edt
}

func GetStudents() []models.StudentJSON {
	data := parseStudentJSON("students.json")
	students := parseStudent(*data)
	return *students
}
