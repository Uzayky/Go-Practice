package Staircase

// https://www.hackerrank.com/challenges/staircase/problem solution.

func StairCase(n int32) string {
	var result string
	// For loop for every iteration
	for i := int32(1); i <= n; i++ {
		// For loop for spaces to be put in ascending steps
		for j := int32(1); j <= n-i; j++ {
			result += " "
		}
		// For loop for "#" character to be put in ascending steps
		for k := int32(1); k <= i; k++ {
			result += "#"
		}
		result += "\n"
	}
	return result
}
