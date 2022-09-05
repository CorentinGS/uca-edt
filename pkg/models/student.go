package models

type StudentJSON struct {
	UUID    string `json:"uuid"`
	Courses map[string]string
}

type StudentEDT struct {
	UUID    string    `json:"uuid"`
	Courses CourseEDT `json:"courses"`
}
