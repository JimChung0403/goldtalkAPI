package queue

import (
	"container/list"
	"sync"
)

const (
	minQSize = 16
)

// Queue defines a queue data structure.
type Queue struct {
	sync.Mutex
	buf  *list.List
}

// New constructor
func New() *Queue {

	return &Queue{
		buf: list.New(),
	}
}

// Enqueue an element.
func (q *Queue) Enqueue(e interface{}) {

	q.Lock()
	defer q.Unlock()
	q.buf.PushBack(e)
}

// Peek a head element without removing
func (q *Queue) Peek() *interface{} {

	if q.buf.Front() == nil {
		return nil
	}

	return &q.buf.Front().Value
}

// Dequeue an element
func (q *Queue) Dequeue() *interface{} {

	q.Lock()
	defer q.Unlock()
	ele := q.buf.Front()
	if ele == nil {
		return nil
	} else {
		q.buf.Remove(ele)
	}
	return &ele.Value
}

// Size returns the length of queue.
func (q *Queue) Size() int {
	q.Lock()
	defer q.Unlock()
	return q.buf.Len()
}
