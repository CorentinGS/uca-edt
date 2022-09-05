package models

type StudentJSON struct {
	UUID    string `json:"uuid"`
	Courses map[string]string
}

type StudentEDT map[string][]CourseEDT

func (s StudentEDT) Print(uuid string) {
	for _, course := range s[uuid] {
		course.Print()
	}
}
