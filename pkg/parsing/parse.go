package parsing

import (
	"github.com/bytedance/sonic"
	"github.com/corentings/uca-edt/pkg/models"
	"log"
	"os"
)

func parseEdtJSON(fileName string) *DataEdtJSON {
	data := new(DataEdtJSON)

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
	data := new(DataStudentJSON)

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
	students := new([]models.StudentJSON)

	for key, value := range data {
		*students = append(*students, models.StudentJSON{
			UUID:    key,
			Courses: value,
		})
	}

	return students
}

func parseEdt(data DataEdtJSON) *CourseEdt {
	edt := CourseEdt{}

	for key, value := range data {
		for key2, value2 := range value {
			for _, course := range value2 {
				if entry, ok := edt[course.Name]; ok {
					entry.CourseEDT = append(entry.CourseEDT, models.CourseEDT{
						Day:      key,
						Hour:     key2,
						Salle:    course.Salle,
						Unparsed: course.Unparsed,
						Groupe:   course.Group,
						Name:     course.Name,
						Type:     course.Type,
					})
					edt[course.Name] = entry
				} else {
					edt[course.Name] = models.CourseData{
						Name: course.Name,
						CourseEDT: []models.CourseEDT{{
							Day:      key,
							Hour:     key2,
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
