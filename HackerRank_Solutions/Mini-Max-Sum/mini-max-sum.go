package Mini_Max_Sum

func MiniMaxSum(arr [5]int32) (int, int) {
	minValue := arr[0]
	maxValue := arr[0]
	var totalSum int32 = 0
	for _, v := range arr {
		if v < minValue {
			minValue = v
		} else {
			maxValue = v
		}
		totalSum += v
	}
	maxSum := totalSum - minValue
	minSum := totalSum - maxValue

	return int(minSum), int(maxSum)
}
