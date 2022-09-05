package models

import (
	"fmt"
	"strings"
)

type Course struct {
	Name     string   `json:"name"`
	Salle    string   `json:"salle"`
	Group    []string `json:"group"`
	Unparsed string   `json:"unparsed"`
}

type CourseData struct {
	Name      string      `json:"name"`
	CourseEDT []CourseEDT `json:"edt"`
}

func (c CourseData) String() {
	fmt.Printf("Course: %s\n\n", c.Name)
	for _, course := range c.CourseEDT {
		unparsed := strings.ReplaceAll(course.Unparsed, "\n", " ")
		unparsed = strings.ReplaceAll(unparsed, "\r", " ")
		fmt.Printf("%s - %s\nSalle: %s\nGroupe:%s\nUnparsed: %s\n", course.Day, course.Hour, course.Salle, course.Groupe, unparsed)
		fmt.Println("--------------------------------------------------")
	}
}

type CourseEDT struct {
	Day      string   `json:"day"`
	Hour     string   `json:"hour"`
	Salle    string   `json:"salle"`
	Unparsed string   `json:"unparsed"`
	Groupe   []string `json:"groupe"`
}
