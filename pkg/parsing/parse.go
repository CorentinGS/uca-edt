package parsing

import (
	"github.com/bytedance/sonic"
	"github.com/corentings/uca-edt/pkg/models"
	"log"
	"os"
)

func parseEdtJSON(fileName string) *DataEdtJSON {
	data := new(DataEdtJSON) // Create a new DataEdtJSON

	// Open the file
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	// Read the file
	err = sonic.Unmarshal(file, &data)
	if err != nil {
		log.Panic(err)
	}

	return data
}

func parseStudentJSON(fileName string) *DataStudentJSON {
	data := new(DataStudentJSON) // Create a new DataStudentJSON

	// Open the file
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	// Read the file
	err = sonic.Unmarshal(file, &data)
	if err != nil {
		log.Panic(err)
	}

	return data
}

func parseStudent(data DataStudentJSON) *[]models.StudentJSON {
	students := new([]models.StudentJSON) // Create a new slice of StudentJSON

	// for each student
	for key, value := range data {
		// Create a new StudentJSON and append it to the slice
		*students = append(*students, models.StudentJSON{
			UUID:    key,
			Courses: value,
		})
	}

	return students
}

func parseEdt(data DataEdtJSON) *CourseEdt {
	edt := CourseEdt{} // Create a new CourseEdt

	// for each day
	for day, value := range data {
		// for each hour
		for hour, value2 := range value {
			// for each course
			for _, course := range value2 {
				// If the course is in the edt
				if entry, ok := edt[course.Name]; ok {
					// Append the course to the edt
					entry.CourseEDT = append(entry.CourseEDT, models.CourseEDT{
						Day:      day,
						Hour:     hour,
						Salle:    course.Salle,
						Unparsed: course.Unparsed,
						Groupe:   course.Group,
						Name:     course.Name,
						Type:     course.Type,
					})
					// Update the edt using the new entry
					edt[course.Name] = entry
				} else {
					// Create a new CourseData and add it to the edt
					edt[course.Name] = models.CourseData{
						Name: course.Name,
						CourseEDT: []models.CourseEDT{{
							Day:      day,
							Hour:     hour,
							Salle:    course.Salle,
							Unparsed: course.Unparsed,
							Groupe:   course.Group,
							Name:     course.Name,
							Type:     course.Type,
						},
						},
					}
				}
			}
		}
	}
	return &edt
}
