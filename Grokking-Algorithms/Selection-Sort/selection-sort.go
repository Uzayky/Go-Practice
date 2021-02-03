package Selection_Sort

func SelectionSort(arr []int32) []int32 {

	var newArray []int32
	for len(arr) > 0 {
		smallestElementIndex := FindSmallestElementIndex(arr)
		newArray = append(newArray, arr[smallestElementIndex])
		arr = append(arr[:smallestElementIndex], arr[smallestElementIndex+1:]...)
	}
	return newArray
}

func FindSmallestElementIndex(arr []int32) int32 {

	smallestElement := arr[0]
	smallestElementIndex := 0
	for i, v := range arr {
		if arr[i] < smallestElement {
			smallestElement = v
			smallestElementIndex = i
		}
	}
	return int32(smallestElementIndex)
}
