package prioritysort

import (
	"testing"

	"github.com/jian-hua-he/go-alg/sort/heapsort"
)

type Item struct {
	Priority int
	Name     string
}

func TestCase1(t *testing.T) {
	arr := []interface{}{
		Item{
			Name:     "Foo",
			Priority: 1,
		},
		Item{
			Name:     "Bar",
			Priority: 2,
		},
		Item{
			Name:     "Fuzz",
			Priority: 1,
		},
		Item{
			Name:     "Buzz",
			Priority: 1,
		},
	}

	result := heapsort.Sort(arr, func(a interface{}, b interface{}) bool {
		return a.(Item).Priority > b.(Item).Priority
	})

	got1 := result[0].(Item).Name
	want1 := "Foo"
	if got1 != want1 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got1, want1)
	}

	got2 := result[1].(Item).Name
	want2 := "Fuzz"
	if got2 != want2 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got2, want2)
	}

	got3 := result[2].(Item).Name
	want3 := "Buzz"
	if got3 != want3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got3, want3)
	}

	got4 := result[3].(Item).Name
	want4 := "Bar"
	if got4 != want4 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got4, want4)
	}
}
