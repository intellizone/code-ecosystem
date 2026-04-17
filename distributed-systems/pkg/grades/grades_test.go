package grades

import (
	"fmt"
	"testing"
)

var student1 = Student{
	ID:        1,
	FirstName: "Suriya Prakash",
	LastName:  "M",
	Grades: []Grade{
		{
			Title: "Quiz 1",
			Type:  GradeQuiz,
			Score: 95,
		},
		{
			Title: "Quiz 2",
			Type:  GradeQuiz,
			Score: 100,
		},
		{
			Title: "Quiz 3",
			Type:  GradeQuiz,
			Score: 90,
		},
	},
}

func TestAverage(t *testing.T) {
	fmt.Println(student1.Average())
}
