package heapsort

import "testing"

func TestCase1(t *testing.T) {
	input := []interface{}{9, 7, 8, 6, 4, 3, 2, 5, 1}
	got := Sort(input, func(a interface{}, b interface{}) bool {
		return a.(int) > b.(int)
	})
	want := []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}

	if len(got) != len(want) {
		t.Errorf("Length of sort result should equal excepted result, got %v, want %v", len(got), len(want))
	}

	for i := 0; i < len(got); i += 1 {
		if got[i] != want[i] {
			t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got[i], want[i])
		}
	}
}
