package parsing

import "github.com/corentings/uca-edt/pkg/models"

type CourseEdt map[string]models.CourseData
type DataEdtJSON map[string]map[string][]models.Course
type DataStudentJSON map[string]map[string]string

const (
	EdtJSONFile     = "parsed.json"
	StudentJSONFile = "parsedGroupe.json"
)

func GetCourseEdt() CourseEdt {
	data := parseEdtJSON(EdtJSONFile)
	edt := parseEdt(*data)
	return *edt
}

func GetStudents() []models.StudentJSON {
	data := parseStudentJSON(StudentJSONFile)
	students := parseStudent(*data)
	return *students
}
