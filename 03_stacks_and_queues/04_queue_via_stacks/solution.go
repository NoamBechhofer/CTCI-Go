package queueviastacks

import (
	"sync"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type QueueViaStacks[T any] struct {
	mu sync.RWMutex

	addStack    lib.ArrayStack[T]
	removeStack lib.ArrayStack[T]
}

// caller is responsible for holding a write lock before calling
func (queue *QueueViaStacks[T]) shuntIfNecessary() {
	if !queue.removeStack.IsEmpty() || queue.addStack.IsEmpty() {
		return
	}

	for !queue.addStack.IsEmpty() {
		pop, _ := queue.addStack.Pop()
		queue.removeStack.Push(pop)
	}
}

func (queue *QueueViaStacks[T]) Add(ele T) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.addStack.Push(ele)
}

func (queue *QueueViaStacks[T]) IsEmpty() bool {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return queue.addStack.IsEmpty() && queue.removeStack.IsEmpty()
}

func (queue *QueueViaStacks[T]) Peek() (T, bool) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.shuntIfNecessary()

	if queue.removeStack.IsEmpty() {
		var zero T
		return zero, false
	}

	return queue.removeStack.Peek()
}

func (queue *QueueViaStacks[T]) Remove() (T, bool) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.shuntIfNecessary()

	if queue.removeStack.IsEmpty() {
		var zero T
		return zero, false
	}

	return queue.removeStack.Pop()

}

func (queue *QueueViaStacks[T]) Size() int {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return queue.addStack.Size() + queue.removeStack.Size()
}
