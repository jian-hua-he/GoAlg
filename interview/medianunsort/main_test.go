package medianunsort

import "testing"

func TestCast1(t *testing.T) {
	input := []float64{1, 3, 3, 6, 7, 8, 9}
	want := 6.0
	got := FindMedian(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got \"%v\", want \"%v\"", input, got, want)
	}
}

func TestCast2(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 8, 9}
	want := 4.5
	got := FindMedian(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got \"%v\", want \"%v\"", input, got, want)
	}
}
