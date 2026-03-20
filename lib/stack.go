package lib

import "sync"

type Stack[T any] struct {
	mu    sync.RWMutex
	slice []T
}

func (stack *Stack[T]) IsEmpty() bool {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	return len(stack.slice) == 0
}

func (stack *Stack[T]) Peek() (T, bool) {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	if len(stack.slice) == 0 {
		var zero T
		return zero, false
	}
	return stack.slice[len(stack.slice)-1], true
}

func (stack *Stack[T]) Pop() (T, bool) {
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

func (stack *Stack[T]) Push(ele T) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	stack.slice = append(stack.slice, ele)
}
