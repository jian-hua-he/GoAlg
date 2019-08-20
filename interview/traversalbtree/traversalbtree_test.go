package traversalbtree

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	bt := &BTree{Value: 5}
	bt.Left = &BTree{Value: 2}
	bt.Right = &BTree{Value: 3}
	bt.Left.Left = &BTree{Value: 20}
	bt.Left.Right = &BTree{Value: 21}
	bt.Right.Left = &BTree{Value: 10}

	got := solution(bt)
	want := 4
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase2(t *testing.T) {
	bt := &BTree{Value: 8}
	bt.Left = &BTree{Value: 3}
	bt.Right = &BTree{Value: 2}
	bt.Left.Left = &BTree{Value: 8}

	got := solution(bt)
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase3(t *testing.T) {
	bt := &BTree{Value: 2}
	got := solution(bt)
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}

func TestCase4(t *testing.T) {
	var bt *BTree = nil
	got := solution(bt)
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Data was incorrect, got %#v, want %#v", got, want)
	}
}
