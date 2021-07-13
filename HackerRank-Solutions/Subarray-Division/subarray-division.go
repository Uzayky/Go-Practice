package Subarray_Division

// https://www.hackerrank.com/challenges/the-birthday-bar/problem

func SubarrayDivision(chocolateBar []int32, birthDay int32, birthMonth int32) int32 {
	var sum, count int32
	if birthMonth > int32(len(chocolateBar)) {
		return 0
	}
	for i := 0; int32(i) < birthMonth; i++ {
		sum += chocolateBar[i]
	}
	if sum == birthDay {
		count++
	}
	subArray := append(chocolateBar[1:])
	return count + SubarrayDivision(subArray, birthDay, birthMonth)
}
