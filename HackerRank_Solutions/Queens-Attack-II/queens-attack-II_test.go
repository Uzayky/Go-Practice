package Queens_Attack_II

import (
	"testing"
)

func TestQueensAttack(t *testing.T) {

	obstacles := [][]int32{{5, 5}, {4, 2}, {2, 3}}

	t.Run("collection of any size", func(t *testing.T) {

		got := QueensAttack(5, 3, 4, 3, obstacles)
		expected := int32(10)

		if got != expected {
			t.Errorf("got %d expected %d of given obstacle slice %v", got, expected, obstacles)
		}
	})
}
