package core

import (
	"fmt"
	"github.com/corentings/uca-edt/pkg/models"
	"github.com/corentings/uca-edt/pkg/parsing"
	"sync"
)

var courseEdt parsing.CourseEdt

func ComputeStudentEDT() {
	students := []models.StudentJSON{
		{
			UUID: "1",
			Courses: map[string]string{
				"MMAG":   "2",
				"MF":     "1",
				"Chimie": "1",
			},
		},
		{
			UUID: "2",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "3",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "4",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "5",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "6",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "7",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
		{
			UUID: "8",
			Courses: map[string]string{
				"MMAG":   "1",
				"MF":     "2",
				"Chimie": "2",
			},
		},
	}

	edt := CreateStudentsEDT(students)
	fmt.Println(len(edt))
	edt.Print("1")
}

func ComputeStudent(student models.StudentJSON) []models.CourseEDT {
	var courses []models.CourseEDT

	for key, value := range student.Courses {
		for _, course := range courseEdt[key].CourseEDT {
			if course.Groupe == value || course.Type == "CM" {
				courses = append(courses, course)
			}
		}
	}

	return courses
}

func CreateStudentsEDT(students []models.StudentJSON) models.StudentEDT {
	courseEdt = parsing.GetCourseEdt()

	type studentEDTChan struct {
		UUID string `json:"uuid"`
		EDT  []models.CourseEDT
	}

	edt := make(models.StudentEDT, len(students))

	wg := sync.WaitGroup{}
	workers := 4
	wg.Add(workers)
	M := len(students) / workers

	ch := make(chan studentEDTChan, len(students))
	var mutex = &sync.Mutex{}

	for i := 0; i < workers; i++ {
		hi, lo := i*M, (i+1)*M
		if i == workers-1 {
			lo = len(students)
		}
		subStudents := students[hi:lo]
		go func() {
			for _, student := range subStudents {
				studentEDT := studentEDTChan{
					UUID: student.UUID,
					EDT:  ComputeStudent(student),
				}
				mutex.Lock()
				ch <- studentEDT
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	for student := range ch {
		edt[student.UUID] = student.EDT
	}

	return edt

}
