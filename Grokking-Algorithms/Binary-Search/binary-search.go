package Binary_Search

import "sort"

func BinarySearch(arr []int32, value int32) int32 {
	SortSlice(arr)
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == value {
			return int32(mid)
		} else if guess < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func SortSlice(arr []int32) []int32 {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
