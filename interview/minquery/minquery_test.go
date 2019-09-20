package minquery

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	keywords := []string{"hello", "happy", "great"}
	maxLen := 50
	want := []string{"hello OR happy OR great"}
	got := generateQueries(keywords, maxLen)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Data was incorrect, keywords %#v, max len %#v, got %#v, want %#v", keywords, maxLen, got, want)
	}
}
