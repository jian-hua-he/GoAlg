package queue

import "testing"

func TestQueue(t *testing.T) {

	i1 := NewItem(999)
	i2 := NewItem("String 1")
	i3 := NewItem(20)

	queue := NewQueue()
	if !queue.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.IsEmpty(), true)
	}

	queue.Enqueue(i1)
	queue.Enqueue(i2)
	queue.Enqueue(i3)

	if queue.Size() != 3 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.Size(), 3)
		queue.Print()
	}

	dequeueItem, err := queue.Dequeue()

	if err != nil {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", err, nil)
		queue.Print()
	}
	if dequeueItem.GetValue() != i1.GetValue() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", dequeueItem.GetValue(), i1.GetValue())
		queue.Print()
	}
	if queue.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.IsEmpty(), false)
		queue.Print()
	}
	if queue.Size() != 2 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.Size(), 2)
		queue.Print()
	}

	queue.Clear()
	if queue.Size() != 0 {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.Size(), 0)
		queue.Print()
	}
	if !queue.IsEmpty() {
		t.Errorf("Data was incorrect, got \"%v\", want \"%v\"", queue.IsEmpty(), true)
		queue.Print()
	}

}
