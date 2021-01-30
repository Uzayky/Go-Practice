package Mini_Max_Sum

func MiniMaxSum(arr [5]int32) (int, int) {
	minValue := arr[0]
	maxValue := arr[0]
	var totalSum int64 = 0
	for _, v := range arr {
		if v < minValue {
			minValue = v
		} else if v > maxValue {
			maxValue = v
		}
		totalSum += int64(v)
	}
	maxSum := totalSum - int64(minValue)
	minSum := totalSum - int64(maxValue)

	return int(minSum), int(maxSum)
}
