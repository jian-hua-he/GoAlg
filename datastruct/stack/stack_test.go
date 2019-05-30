package stack

import "testing"

func TestStack(t *testing.T) {

	stack := NewStack()

	if stack.Size() != 0 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", stack.Size(), 0)
		stack.Print()
	}
	if !stack.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", stack.IsEmpty(), true)
		stack.Print()
	}

	i1 := NewItem(777)
	i2 := NewItem("String")
	i3 := NewItem(6.6)
	stack.Push(i1)
	stack.Push(i2)
	stack.Push(i3)

	if stack.Size() != 3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", stack.Size(), 3)
		stack.Print()
	}
	if stack.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", stack.IsEmpty(), false)
		stack.Print()
	}

	poped, err := stack.Pop()
	if err != nil {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", err, nil)
		stack.Print()
	}
	if poped.GetValue() != i3.GetValue() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", poped.GetValue(), i3.GetValue())
		stack.Print()
	}
	if stack.Size() != 2 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", stack.Size(), 2)
		stack.Print()
	}

}
