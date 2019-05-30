package medianunsort

import "testing"

func TestFindMedian(t *testing.T) {
	input1 := []float64{1, 3, 3, 6, 7, 8, 9}
	want1 := 6.0
	got1 := FindMedian(input1)
	if got1 != want1 {
		t.Errorf("Data was incorrect, input %v, got \"%v\", want \"%v\"", input1, got1, want1)
	}

	input2 := []float64{1, 2, 3, 4, 5, 6, 8, 9}
	want2 := 4.5
	got2 := FindMedian(input2)
	if got2 != want2 {
		t.Errorf("Data was incorrect, input %v, got \"%v\", want \"%v\"", input2, got2, want2)
	}
}
