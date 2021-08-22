package queue

import (
	"container/heap"
	"sync"
)

type Queue interface {
	Enqueue(item Item)
	Dequeue() Item
}

type queue struct {
	items []Item
	mu    sync.Mutex
}

func (q *queue) Enqueue(item Item) {
	defer q.synchronise()()
	heap.Push(q, item)
}

func (q *queue) Dequeue() Item {
	defer q.synchronise()()
	item := heap.Pop(q)
	switch item.(type) {
	case Item:
		return item.(Item)
	default:
		return nil
	}
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Less(i, j int) bool {
	return q.items[i].Priority() < q.items[j].Priority()
}

func (q *queue) Swap(i, j int) {

	if i >= len(q.items) {
		return
	}
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *queue) Push(i interface{}) {
	q.items = append(q.items, i.(Item))
}

func (q *queue) Pop() interface{} {

	if len(q.items) == 0 {
		return nil
	}
	minItem := q.items[len(q.items)-1]
	q.items = q.items[0:len(q.items)-1]
	return minItem
}

func (q *queue) synchronise() func() {
	q.mu.Lock()
	return q.mu.Unlock
}

func NewQueue(items []Item) Queue {
	return &queue{items: items}
}
