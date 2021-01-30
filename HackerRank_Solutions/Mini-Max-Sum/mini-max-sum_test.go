package Mini_Max_Sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int32{2, 1, 3, 4, 5}

		minSum, maxSum := MiniMaxSum(numbers)
		expectedMinSum := 10
		expectedMaxSum := 14

		if minSum != expectedMinSum || maxSum != expectedMaxSum {
			t.Errorf("minSum %d maxSum %d expectedMinSum %d expectedMaxSum %d given, %v", minSum, maxSum, expectedMinSum, expectedMaxSum, numbers)
		}
	})
}
