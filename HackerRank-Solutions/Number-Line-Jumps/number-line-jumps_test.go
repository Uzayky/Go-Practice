package Number_Line_Jumps

import (
	"testing"
)

func TestKangaroo(t *testing.T) {

	isKangaroosOnSameSpot := func(t testing.TB, got, expected string) {
		t.Helper()

		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	}

	t.Run("", func(t *testing.T) {

		got := Kangaroo(0, 3, 4, 2)
		expected := "YES"
		isKangaroosOnSameSpot(t, got, expected)
	})

	t.Run("", func(t *testing.T) {

		got := Kangaroo(43, 2, 70, 2)
		expected := "NO"
		isKangaroosOnSameSpot(t, got, expected)

	})
}
