package linkedli

import "testing"

func TestLinkedList(t *testing.T) {

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	l := NewLinkedList()

	if l.Length() != 0 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.Length(), 0)
		l.Print()
	}
	if !l.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.IsEmpty(), true)
		l.Print()
	}

	l.Append(n1)
	l.Append(n2)
	l.Append(n3)

	if l.Length() != 3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.Length(), 3)
		l.Print()
	}
	if l.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.IsEmpty(), false)
		l.Print()
	}
	if l.First().GetValue() != n1.GetValue() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.First().GetValue(), n1.GetValue())
	}

	n4 := NewNode(4)
	l.Insert(1, n4)

	if l.Length() != 4 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.Length(), 4)
		l.Print()
	}
	if l.First().Next().GetValue() != n4.GetValue() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.First().Next().GetValue(), n4.GetValue())
		l.Print()
	}

	l.RemoveAt(1)
	if l.Length() != 3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.Length(), 3)
		l.Print()
	}
	if l.First().Next().GetValue() != n2.GetValue() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", l.First().Next().GetValue(), n2.GetValue())
		l.Print()
	}
}
