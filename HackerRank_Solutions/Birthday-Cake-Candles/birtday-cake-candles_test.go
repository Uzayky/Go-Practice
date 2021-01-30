package Birthday_Cake_Candles

import "testing"

func TestBirthDayCakeCandles(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		candles := []int32{4, 4, 1, 2, 6, 5, 6}

		got := birthdayCakeCandles(candles)
		expected := int32(2)

		if got != expected {
			t.Errorf("got %d expected %d of given %v", got, expected, candles)
		}
	})
}
