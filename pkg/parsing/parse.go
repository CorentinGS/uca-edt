package parsing

import (
	"encoding/json"
	"github.com/corentings/uca-edt/pkg/models"
	"log"
	"os"
)

func Parse() {
	data := parseJson()

	edt := parseEdt(*data)

	edt["MMAG"].String()
}

func parseJson() *map[string]map[string][]models.Course {
	data := new(map[string]map[string][]models.Course)

	// Open the file
	file, err := os.ReadFile("parsed.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read the file
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func parseEdt(data map[string]map[string][]models.Course) map[string]models.CourseData {
	edt := map[string]models.CourseData{}

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
						},
						},
					}
				}
			}
		}
	}
	return edt
}
