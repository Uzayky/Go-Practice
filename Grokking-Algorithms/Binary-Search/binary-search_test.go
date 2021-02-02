package Binary_Search

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	arr := []int32{1, 72, 36, 98, 6, 15, 17}

	t.Run("collection of any size", func(t *testing.T) {
		value := 36

		got := BinarySearch(arr, int32(value))
		expected := int32(4)

		if got != expected {
			t.Errorf("got %d expected %d of given %v", got, expected, arr)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {

		got := SortSlice(arr)
		expected := []int32{1, 6, 15, 17, 36, 72, 98}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}
