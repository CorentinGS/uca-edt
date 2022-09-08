package parsing

import "github.com/corentings/uca-edt/pkg/models"

// CourseEdt is a type alias for a map of courses edt
type CourseEdt map[string]models.CourseData

// DataEdtJSON is a type alias for the JSON structure of the edt
type DataEdtJSON map[string]map[string][]models.Course

// DataStudentJSON is a type alias for the JSON structure of the students
type DataStudentJSON map[string]map[string]string

// GetCourseEdt returns a map of students edt
func GetCourseEdt(EdtJSONFile string) CourseEdt {
	data := parseEdtJSON(EdtJSONFile) // Parse the JSON file
	edt := parseEdt(*data)            // Parse the data
	return *edt
}

// GetStudents returns a slice of students
func GetStudents(StudentJSONFile string) []models.StudentJSON {
	data := parseStudentJSON(StudentJSONFile) // Parse the JSON file
	students := parseStudent(*data)           // Parse the data
	return *students
}
