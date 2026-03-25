package stackmin

import (
	"cmp"
	"sync"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type StackMin[T cmp.Ordered] struct {
	mu       sync.RWMutex
	data     lib.ArrayStack[T]
	minStack lib.ArrayStack[T]
}

func (stack *StackMin[T]) IsEmpty() bool {
	stack.mu.RLock()
	defer stack.mu.RUnlock()
	return stack.data.IsEmpty()
}

func (stack *StackMin[T]) Peek() (T, bool) {
	stack.mu.RLock()
	defer stack.mu.RUnlock()
	return stack.data.Peek()
}

func (stack *StackMin[T]) Pop() (T, bool) {
	stack.mu.Lock()
	defer stack.mu.Unlock()
	retVal, retOk := stack.data.Pop()
	minVal, _ := stack.minStack.Peek()
	if retOk && retVal == minVal {
		stack.minStack.Pop()
	}
	return retVal, retOk
}

func (stack *StackMin[T]) Push(ele T) {
	stack.mu.Lock()
	defer stack.mu.Unlock()
	minVal, notEmpty := stack.minStack.Peek()
	if !notEmpty || ele <= minVal {
		stack.minStack.Push(ele)
	}
	stack.data.Push(ele)
}

func (stack *StackMin[T]) Size() int {
	stack.mu.RLock()
	defer stack.mu.RUnlock()
	return stack.data.Size()
}

func (stack *StackMin[T]) Min() (T, bool) {
	stack.mu.RLock()
	defer stack.mu.RUnlock()
	return stack.minStack.Peek()
}
