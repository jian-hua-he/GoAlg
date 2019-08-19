package operatorstr

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "13 DUP 13 7 + 9 POP -"
	want := 7
	got := solution(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase2(t *testing.T) {
	input := " 13 DUP 13 7 + 9 POP -"
	want := -1
	got := solution(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase3(t *testing.T) {
	input := "1 22 3 5 10"
	want := 10
	got := solution(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase4(t *testing.T) {
	input := "a1 22 3 5 *x"
	want := -1
	got := solution(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}
