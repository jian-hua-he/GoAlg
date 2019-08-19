package operatorstr

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"case_1": {
			input: "13 DUP 13 7 + 9 POP -",
			want:  7,
		},
		"case_2": {
			input: " 13 DUP 13 7 + 9 POP -",
			want:  -1,
		},
		"case_3": {
			input: "1 22 3 5 10",
			want:  10,
		},
		"case_4": {
			input: "a1 22 3 5 *x",
			want:  -1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("Data was incorrect, input %#v, got %#v, want %#v", tc.input, got, tc.want)
			}
		})
	}
}

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
