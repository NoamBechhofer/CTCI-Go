package lib

import "sync"

// default implementation of Stack using a slice as backing storage.
type ArrayStack[T any] struct {
	mu    sync.RWMutex
	slice []T
}

func ArrayStackFromSlice[T any](slice []T) *ArrayStack[T] {
	var ret ArrayStack[T]

	for _, ele := range slice {
		ret.Push(ele)
	}

	return &ret
}

func (stack *ArrayStack[T]) ToSlice() []T {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	destination := make([]T, len(stack.slice))
	copy(destination, stack.slice)
	return destination
}

func (stack *ArrayStack[T]) IsEmpty() bool {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	return len(stack.slice) == 0
}

func (stack *ArrayStack[T]) Peek() (T, bool) {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	if len(stack.slice) == 0 {
		var zero T
		return zero, false
	}
	return stack.slice[len(stack.slice)-1], true
}

func (stack *ArrayStack[T]) Pop() (T, bool) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	if len(stack.slice) == 0 {
		var zero T
		return zero, false
	}

	last := len(stack.slice) - 1
	ret := stack.slice[len(stack.slice)-1]

	var zero T
	stack.slice[last] = zero
	stack.slice = stack.slice[:last]

	return ret, true
}

func (stack *ArrayStack[T]) Push(ele T) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	stack.slice = append(stack.slice, ele)
}

func (stack *ArrayStack[T]) Size() int {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	return len(stack.slice)
}
