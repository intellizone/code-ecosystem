package grades

import (
	"math/rand"

	"github.com/go-faker/faker/v4"
)

func init() {
	students = []Student{}
	for i := range 15 {
		students = append(students,
			Student{
				ID:        i + 1,
				FirstName: faker.FirstName(),
				LastName:  faker.LastName(),
				Grades: []Grade{
					{
						Title: "Quiz 1",
						Type:  GradeQuiz,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Week 1 Homework",
						Type:  GradeHomework,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Quiz 2",
						Type:  GradeQuiz,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Week 2 Homework",
						Type:  GradeHomework,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Quiz 3",
						Type:  GradeQuiz,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Week 3 Homework",
						Type:  GradeHomework,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Quiz 4",
						Type:  GradeQuiz,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Week 4 Homework",
						Type:  GradeHomework,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Quiz 5",
						Type:  GradeQuiz,
						Score: rand.Float32() * 100,
					},
					{
						Title: "Week 5 Homework",
						Type:  GradeHomework,
						Score: rand.Float32() * 100,
					},
				},
			})
	}
}
