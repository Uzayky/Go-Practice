package Birthday_Cake_Candles

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func birthdayCakeCandles(candles []int32) int32 {
	maxValue := candles[0]
	var candleCount int32 = 0
	for _, v := range candles {
		if v > maxValue {
			maxValue = v
			candleCount = 0
		}
		if v == maxValue {
			candleCount++
		}
	}

	return candleCount
}
