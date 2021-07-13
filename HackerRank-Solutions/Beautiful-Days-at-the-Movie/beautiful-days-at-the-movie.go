package Beautiful_Days_at_the_Movie

// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem

func BeautifulDays(i int32, j int32, k int32) int32 {

	var count int32

	for z := i; z <= j; z++ {
		if (z-ReverseInteger(z))%k == 0 {
			count++
		}
	}

	return count
}

func ReverseInteger(val int32) int32 {

	var reversedInteger int32
	for val != 0 {
		digit := val % 10
		reversedInteger = reversedInteger*10 + digit
		val /= 10
	}

	return reversedInteger
}
