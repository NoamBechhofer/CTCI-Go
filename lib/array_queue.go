package lib

import "sync"

// default implementation of Queue using a slice as backing storage.
type ArrayQueue[T any] struct {
	mu    sync.RWMutex
	slice []T
}

func (queue *ArrayQueue[T]) Add(ele T) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.slice = append(queue.slice, ele)
}

func (queue *ArrayQueue[T]) IsEmpty() bool {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return len(queue.slice) == 0
}

func (queue *ArrayQueue[T]) Peek() (T, bool) {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	if len(queue.slice) == 0 {
		var zero T
		return zero, false
	}

	return queue.slice[0], true
}

func (queue *ArrayQueue[T]) Remove() (T, bool) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	var zero T

	if len(queue.slice) == 0 {
		return zero, false
	}

	ret := queue.slice[0]

	queue.slice[0] = zero
	queue.slice = queue.slice[1:]

	return ret, true
}

func (queue *ArrayQueue[T]) Size() int {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return len(queue.slice)
}
