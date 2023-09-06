package core

import (
	"sync"

	"github.com/corentings/uca-edt/pkg/database"
	"github.com/corentings/uca-edt/pkg/models"
	"github.com/corentings/uca-edt/pkg/parsing"
)

// ComputeStudentEDT function computes the edt of all the students
func ComputeStudentEDT(studentFile string, edtFile string) {
	students := parsing.GetStudents(studentFile) // Get the students
	courseEdt := parsing.GetCourseEdt(edtFile)   // Get the parsed edt

	edt := CreateStudentsEDT(students, courseEdt) // Create the edt of all the students

	database.StoreEdt(edt) // Store the edt in the database
}

// ComputeStudent function computes the edt of a student
func ComputeStudent(student models.StudentJSON, courseEdt *parsing.CourseEdt) []models.CourseEDT {
	var courses []models.CourseEDT // courses is the list of the courses of the student

	// For each course of the student
	for key, value := range student.Courses {
		// Get the course data and classes
		for _, class := range (*courseEdt)[key].CourseEDT {
			// If the class group is the same as the student's class group or if it's a CM
			if class.Groupe == value || class.Type == "CM" && class.Groupe == "" || class.Groupe == "0" {
				// Add the class to the list of the courses of the student
				courses = append(courses, class)
			}
		}
	}
	return courses // Return the list of the courses of the student
}

// CreateStudentsEDT function creates the edt of a student
func CreateStudentsEDT(students []models.StudentJSON, courseEdt parsing.CourseEdt) models.StudentEDT {
	type studentEDTChan struct {
		UUID string             `json:"uuid"` // UUID is the uuid of the student
		EDT  []models.CourseEDT `json:"edt"`  // EDT is the edt of the student
	}

	edt := make(models.StudentEDT, len(students)) // edt is the edt of all the students

	// Work in parallel
	wg := sync.WaitGroup{}       // wg is a WaitGroup
	workers := 4                 // workers is the number of workers
	wg.Add(workers)              // Add the number of workers to the WaitGroup
	M := len(students) / workers // M is the number of students per worker

	ch := make(chan studentEDTChan, len(students)) // Create a channel
	mutex := &sync.Mutex{}                         // mutex is a mutex to prevent data racing

	// For each worker
	for i := 0; i < workers; i++ {
		hi, lo := i*M, (i+1)*M // hi and lo are the indexes of the students for the worker
		// If it's the last worker
		if i == workers-1 {
			lo = len(students) // Set lo to the number of students
		}
		subStudents := students[hi:lo] // subStudents is the sublist of the students for the worker

		// Create a goroutine
		go func() {
			// For each student
			for _, student := range subStudents {
				// Compute the edt of the student
				studentEDT := studentEDTChan{
					UUID: student.UUID,
					EDT:  ComputeStudent(student, &courseEdt),
				}
				mutex.Lock()     // Lock the mutex
				ch <- studentEDT // Send the edt of the student to the channel
				mutex.Unlock()   // Unlock the mutex
			}
			wg.Done() // Decrement the WaitGroup
		}()
	}
	wg.Wait() // Wait for all the workers to finish
	close(ch) // Close the channel
	// For student in the channel
	for student := range ch {
		edt[student.UUID] = student.EDT // Add the edt of the student to the edt of all the students
	}

	return edt // Return the edt of all the students
}
