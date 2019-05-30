package linkedli

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func NewNode(v int) *Node {
	return &Node{
		value: v,
		next:  nil,
	}
}

func (n Node) GetValue() int {
	return n.value
}

func (n Node) Next() *Node {
	return n.next
}

func (n *Node) UnlinkNext() {
	n.next = nil
}

func (self *Node) AddNode(n *Node) {
	self.next = n
}

type LinkedList struct {
	length int
	list   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		length: 0,
		list:   nil,
	}
}

func (l LinkedList) Length() int {
	return l.length
}

func (l LinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l LinkedList) First() *Node {
	return l.list
}

func (l LinkedList) Last() *Node {
	result := l.list

	for n := l.list; n != nil; n = n.Next() {
		result = n
	}

	return result
}

func (l *LinkedList) Append(n *Node) {
	if l.length == 0 {
		l.list = n
	} else {
		l.Last().AddNode(n)
	}

	l.length += 1
}

func (l *LinkedList) Insert(index int, n *Node) error {
	if index >= l.Length() || index < 0 {
		return errors.New("Invalid index")
	}

	if index == 0 {
		n.AddNode(l.list)
		l.list = n
		l.length += 1

		return nil
	}

	i := 0
	curr := l.list
	var prev *Node
	for i < index {
		prev = curr
		curr = l.list.Next()
		i += 1
	}

	n.AddNode(curr)
	prev.UnlinkNext()
	prev.AddNode(n)
	l.length += 1

	return nil
}

func (l *LinkedList) RemoveAt(index int) error {
	if index >= l.Length() || index < 0 {
		return errors.New("Invalid index")
	}

	if index == 0 {
		l.list = l.list.Next()
		l.length -= 1
		return nil
	}

	i := 0
	curr := l.list
	var prev *Node
	for i < index {
		prev = curr
		curr = curr.Next()
		i += 1
	}

	prev.UnlinkNext()
	prev.AddNode(curr.Next())
	l.length -= 1

	return nil
}

func (l LinkedList) Print() {
	for n := l.list; n != nil; n = n.Next() {
		fmt.Println(n.GetValue())
	}
}
