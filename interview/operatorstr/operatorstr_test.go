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
