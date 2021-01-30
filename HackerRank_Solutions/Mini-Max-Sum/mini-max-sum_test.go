package Mini_Max_Sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int32{156873294, 719583602, 581240736, 605827319, 895647130}

		minSum, maxSum := MiniMaxSum(numbers)
		expectedMinSum := 2063524951
		expectedMaxSum := 2802298787

		if minSum != expectedMinSum || maxSum != expectedMaxSum {
			t.Errorf("minSum %d maxSum %d expectedMinSum %d expectedMaxSum %d given, %v", minSum, maxSum, expectedMinSum, expectedMaxSum, numbers)
		}
	})
}
