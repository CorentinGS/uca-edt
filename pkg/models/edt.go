package models

import (
	"fmt"
	"strings"
)

type Course struct {
	Name     string `json:"name"`
	Salle    string `json:"salle"`
	Group    string `json:"group"`
	Type     string `json:"type"`
	Unparsed string `json:"unparsed"`
}

type CourseData struct {
	Name      string      `json:"name"`
	CourseEDT []CourseEDT `json:"edt"`
}

func (c CourseData) Print() {
	fmt.Printf("Course: %s\n\n", c.Name)
	for _, course := range c.CourseEDT {
		unparsed := strings.ReplaceAll(course.Unparsed, "\n", " ")
		unparsed = strings.ReplaceAll(unparsed, "\r", " ")
		fmt.Printf("%s - %s\nSalle: %s\nGroupe:%s\nUnparsed: %s\n", course.Day, course.Hour, course.Salle, course.Groupe, unparsed)
		fmt.Println("--------------------------------------------------")
	}
}

type CourseEDT struct {
	Day      string `json:"day" bson:"day"`
	Hour     string `json:"hour" bson:"hour"`
	Salle    string `json:"salle" bson:"salle"`
	Unparsed string `json:"unparsed" bson:"unparsed"`
	Groupe   string `json:"groupe" bson:"groupe"`
	Type     string `json:"type" bson:"type"`
	Name     string `json:"name" bson:"name"`
}

func (c CourseEDT) Print() {
	fmt.Printf("Course: %s\n", c.Name)
	unparsed := strings.ReplaceAll(c.Unparsed, "\n", " ")
	unparsed = strings.ReplaceAll(unparsed, "\r", " ")
	fmt.Printf("%s - %s\nSalle: %s\nGroupe:%s\nUnparsed: %s\n", c.Day, c.Hour, c.Salle, c.Groupe, unparsed)
	fmt.Println("--------------------------------------------------")
}
