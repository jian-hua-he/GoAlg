package minquery

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	tests := map[string]struct {
		keywords []string
		max      int
		want     int
	}{
		"case_1": {
			keywords: []string{"hello", "happy", "great"},
			max:      50,
			want:     1,
		},
		"case_2": {
			keywords: []string{"foo", "fuzz", "bar", "buzz"},
			max:      10,
			want:     3,
		},
		"case_3": {
			keywords: []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiii", "jjjjjj", "kkkkk", "llll", "mmm", "nn", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
			max:      13,
			want:     11,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := len(generateQueries(tc.keywords, tc.max))
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("Data was incorrect, keywords %#v, max len %#v, got %#v, want %#v", tc.keywords, tc.max, got, tc.want)
			}
		})
	}
}
