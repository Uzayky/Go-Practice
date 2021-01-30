package Grading_Students

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {

	t.Run("", func(t *testing.T) {

		got := GradingStudents([]int32{73, 67, 38, 33})
		expected := []int32{75, 67, 40, 33}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}
