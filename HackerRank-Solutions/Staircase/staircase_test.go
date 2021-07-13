package Staircase

import "testing"

func TestStairCase(t *testing.T) {
	repeated := StairCase(2)
	expected := " #\n##\n"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StairCase(int32(i))
	}
}
