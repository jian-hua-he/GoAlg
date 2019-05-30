package stack

import (
	"errors"
	"fmt"
)

type Item struct {
	value interface{}
}

func NewItem(i interface{}) Item {
	return Item{i}
}

func (i Item) GetValue() interface{} {
	return i.value
}

type Stack struct {
	items []Item
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(i Item) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (Item, error) {
	size := len(s.items)
	if size == 0 {
		return Item{nil}, errors.New("Try to pop empty stack")
	}

	result := s.items[size-1]
	s.items = append(s.items[0 : size-1])

	return result, nil
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Stack) Size() int {
	return len(s.items)
}

func (s Stack) Print() {
	for _, i := range s.items {
		fmt.Println(i.GetValue())
	}
}

func main() {
	i1 := NewItem(777)
	i2 := NewItem("String")

	stack := NewStack()
	fmt.Printf("Stack size is: %+v\n", stack.Size())
	fmt.Printf("Is stack empty: %+v\n", stack.IsEmpty())

	stack.Push(i1)
	stack.Push(i2)
	fmt.Printf("Stack size are: %+v\n", stack.Size())
	fmt.Printf("Is stack empty: %+v\n", stack.IsEmpty())
	fmt.Println("===== Stack Content =====")
	stack.Print()
	fmt.Println("=========================")

	i3 := NewItem(6.6)
	stack.Push(i3)
	fmt.Println("===== Stack Content =====")
	stack.Print()
	fmt.Println("=========================")

	poped, _ := stack.Pop()
	fmt.Printf("Poped item is: %+v\n", poped.GetValue())
	fmt.Printf("Stack size is: %+v\n", stack.Size())
	fmt.Println("===== Stack Content =====")
	stack.Print()
	fmt.Println("=========================")
}
