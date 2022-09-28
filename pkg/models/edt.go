package models

import (
	"fmt"
	"strings"
)

// Course struct
type Course struct {
	Name     string `json:"name"`     // Name of the course
	Salle    string `json:"salle"`    // Room of the course
	Group    string `json:"group"`    // Group of the course
	Type     string `json:"type"`     // Type of the course
	Unparsed string `json:"unparsed"` // Unparsed data of the course
}

// CourseData struct
type CourseData struct {
	Name      string      `json:"name" bson:"name"`      // Name of the course
	CourseEDT []CourseEDT `json:"edt" bson:"course_edt"` // Course's EDT
}

// Print prints the course's EDT
func (c CourseData) Print() {
	fmt.Printf("Course: %s\n\n", c.Name) // Print the name of the course
	// for each course in the course's EDT
	for _, course := range c.CourseEDT {
		// Clean the unparsed data
		unparsed := strings.ReplaceAll(course.Unparsed, "\n", " ")
		unparsed = strings.ReplaceAll(unparsed, "\r", " ")
		// Print the course's EDT
		fmt.Printf("%s - %s\nSalle: %s\nGroupe:%s\nUnparsed: %s\n", course.Day, course.Hour, course.Salle, course.Groupe, unparsed)
		fmt.Println("--------------------------------------------------") // Print a separator
	}
}

// CourseEDT struct
type CourseEDT struct {
	Day      string `json:"day" bson:"day"`           // Day of the course
	Hour     string `json:"hour" bson:"hour"`         // Hour of the course
	Salle    string `json:"salle" bson:"salle"`       // Room of the course
	Unparsed string `json:"unparsed" bson:"unparsed"` // Unparsed data of the course
	Groupe   string `json:"groupe" bson:"groupe"`     // Group of the course
	Type     string `json:"type" bson:"type"`         // Type of the course
	Name     string `json:"name" bson:"name"`         // Name of the course
}

// Print prints the course's EDT
func (c CourseEDT) Print() {
	fmt.Printf("Course: %s\n", c.Name)
	unparsed := strings.ReplaceAll(c.Unparsed, "\n", " ")
	unparsed = strings.ReplaceAll(unparsed, "\r", " ")
	fmt.Printf("%s - %s\nSalle: %s\nGroupe:%s\nUnparsed: %s\n", c.Day, c.Hour, c.Salle, c.Groupe, unparsed)
	fmt.Println("--------------------------------------------------")
}
