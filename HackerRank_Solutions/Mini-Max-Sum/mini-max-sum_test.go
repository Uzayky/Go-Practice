package Mini_Max_Sum

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{2, 1, 3, 4, 5}

	minSum, maxSum := MiniMaxSum(numbers)
	expectedMinSum := 14
	expectedMaxSum := 10

	if minSum != expectedMinSum || maxSum != expectedMaxSum {
		t.Errorf("minSum %d maxSum %d expectedMinSum %d expectedMaxSum %d given, %v", minSum, maxSum, expectedMinSum, expectedMaxSum, numbers)
	}
}
