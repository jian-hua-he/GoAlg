package mergesort

import "testing"

func TestSortCase1(t *testing.T) {
	input := []interface{}{7, 1, 8, 6, 4, 3, 2, 5, 9}
	got := Sort(input, func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
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

func TestSortCase2(t *testing.T) {
	input := []interface{}{1, 5, 3, 2, 4}
	got := Sort(input, func(a interface{}, b interface{}) bool {
		return a.(int) > b.(int)
	})
	want := []interface{}{1, 2, 3, 4, 5}

	if len(got) != len(want) {
		t.Errorf("Length of sort result should equal excepted result, got %v, want %v", len(got), len(want))
	}

	for i := 0; i < len(got); i += 1 {
		if got[i] != want[i] {
			t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got[i], want[i])
		}
	}
}

func TestMergeCase1(t *testing.T) {
	left := []interface{}{1, 3, 4, 7}
	right := []interface{}{2, 5, 6}
	got := merge(left, right, func(l interface{}, r interface{}) bool {
		return l.(int) > r.(int)
	})

	want := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(want); i += 1 {
		if got[i] != want[i] {
			t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got[i], want[i])
		}
	}
}

func TestMergeCase2(t *testing.T) {
	left := []interface{}{4, 2}
	right := []interface{}{8, 7, 6, 5, 3, 1}
	got := merge(left, right, func(l interface{}, r interface{}) bool {
		return l.(int) < r.(int)
	})

	want := []int{8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(want); i += 1 {
		if got[i] != want[i] {
			t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got[i], want[i])
		}
	}
}
