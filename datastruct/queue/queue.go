package queue

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

type Queue struct {
	items []Item
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(i Item) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() (Item, error) {
	if len(q.items) == 0 {
		return Item{nil}, errors.New("Try to dequeue empty subsequence")
	}

	result := q.items[0]
	q.items = q.items[1:]

	return result, nil
}

func (q *Queue) Clear() {
	q.items = make([]Item, 0)
}

func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q Queue) Size() int {
	return len(q.items)
}

func (q Queue) Print() {
	for _, item := range q.items {
		fmt.Println(item.GetValue())
	}
}
