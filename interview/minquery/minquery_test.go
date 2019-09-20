package minquery

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	keywords := []string{"hello", "happy", "great"}
	maxLen := 50
	want := 1
	got := len(generateQueries(keywords, maxLen))

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Data was incorrect, keywords %#v, max len %#v, got %#v, want %#v", keywords, maxLen, got, want)
	}
}

func TestCase2(t *testing.T) {
	keywords := []string{"foo", "fuzz", "bar", "buzz"}
	maxLen := 10
	want := 3
	got := len(generateQueries(keywords, maxLen))

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Data was incorrect, keywords %#v, max len %#v, got %#v, want %#v", keywords, maxLen, got, want)
	}
}

func TestCase3(t *testing.T) {
	keywords := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiii", "jjjjjj", "kkkkk", "llll", "mmm", "nn", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	maxLen := 13
	want := 11
	got := len(generateQueries(keywords, maxLen))

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Data was incorrect, keywords %#v, max len %#v, got %#v, want %#v", keywords, maxLen, got, want)
	}
}
