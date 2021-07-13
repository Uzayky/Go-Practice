package Grading_Students

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {

	checkGrades := func(t testing.TB, got, want []int32) {
		t.Helper()

		// It's important to note that reflect.DeepEqual is not "type safe", the code will compile even if you did something a bit silly.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {

		got := GradingStudents([]int32{73, 67, 38, 33})
		expected := []int32{75, 67, 40, 33}
		checkGrades(t, got, expected)
	})
}
