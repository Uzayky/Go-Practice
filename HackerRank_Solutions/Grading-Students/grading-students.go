package Grading_Students

// https://www.hackerrank.com/challenges/grading/problem

func GradingStudents(grades []int32) []int32 {
	for idx, v := range grades {
		if 5-(v%5) < 3 && v >= 38 {
			grades[idx] += 5 - (v % 5)
		}
	}

	return grades
}
