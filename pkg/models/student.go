package models

// StudentJSON is the JSON structure of a student
type StudentJSON struct {
	Courses map[string]string
	UUID    string `json:"uuid"`
}

// StudentEDT is an alias type for a map of students edt
type StudentEDT map[string][]CourseEDT

func (s StudentEDT) Print(uuid string) {
	for _, course := range s[uuid] {
		course.Print()
	}
}
