package parsing

import "github.com/corentings/uca-edt/pkg/models"

type CourseEdt map[string]models.CourseData
type DataEdtJSON map[string]map[string][]models.Course
type DataStudentJSON map[string]map[string]string

func GetCourseEdt(EdtJSONFile string) CourseEdt {
	data := parseEdtJSON(EdtJSONFile)
	edt := parseEdt(*data)
	return *edt
}

func GetStudents(StudentJSONFile string) []models.StudentJSON {
	data := parseStudentJSON(StudentJSONFile)
	students := parseStudent(*data)
	return *students
}
