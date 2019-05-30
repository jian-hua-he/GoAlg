package prioritysort

import "testing"

func TestCase1(t *testing.T) {
	arr := []Item{
		{
			Name:     "Foo",
			Priority: 1,
		},
		{
			Name:     "Bar",
			Priority: 2,
		},
		{
			Name:     "Fuzz",
			Priority: 1,
		},
	}

	result := SortPriority(arr)

	got1 := result[0].Name
	want1 := "Foo"
	if got1 != want1 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got1, want1)
	}

	got2 := result[1].Name
	want2 := "Fuzz"
	if got2 != want2 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got2, want2)
	}

	got3 := result[2].Name
	want3 := "Bar"
	if got3 != want3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", got3, want3)
	}
}
