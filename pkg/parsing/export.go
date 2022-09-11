package parsing

import (
	"github.com/corentings/uca-edt/pkg/models"
	"log"
)

// CourseEdt is a type alias for a map of courses edt
type CourseEdt map[string]models.CourseData

// DataEdtJSON is a type alias for the JSON structure of the edt
type DataEdtJSON map[string]map[string][]models.Course

// DataStudentJSON is a type alias for the JSON structure of the students
type DataStudentJSON map[string]map[string]string

// GetCourseEdt returns a map of students edt
func GetCourseEdt(EdtJSONFile string) CourseEdt {
	data := new(DataEdtJSON) // Create a new DataEdtJSON
	// Parse the JSON file
	if err := parseJSON(EdtJSONFile, data); err != nil {
		log.Panicf("Error while parsing edt JSON file: %s", err.Error())
	}
	edt := parseEdt(*data) // Parse the data
	return *edt
}

// GetStudents returns a slice of students
func GetStudents(StudentJSONFile string) []models.StudentJSON {
	data := new(DataStudentJSON) // Create a new DataStudentJSON

	// Parse the JSON file
	if err := parseJSON(StudentJSONFile, data); err != nil {
		log.Panicf("Error while parsing student JSON file: %s", err.Error())
	}
	students := parseStudent(*data) // Parse the data
	return *students
}
