package stackofplates

import (
	"github.com/NoamBechhofer/CTCI-Go/lib"
)

const MAX_STACK_SIZE = 2

type stackOfPlates[T any] struct {
	data *lib.ArrayStack[*lib.ArrayStack[T]]
	top  int
	size uint
}

func NewStackOfPlates[T any]() *stackOfPlates[T] {
	data := lib.ArrayStack[*lib.ArrayStack[T]]{}
	return &stackOfPlates[T]{data: &data, top: 0, size: 0}
}

func (stack *stackOfPlates[T]) IsEmpty() bool {
	return stack.size == 0
}

func (stack *stackOfPlates[T]) Peek() (T, bool) {
	topStack, topStackOk := stack.data.Peek()
	if !topStackOk {
		var zero T
		return zero, topStackOk
	}
	return topStack.Peek()
}

func (stack *stackOfPlates[T]) Pop() (T, bool) {
	if stack.IsEmpty() {
		var zero T
		return zero, false
	}

	topStack, _ := stack.data.Peek()
	retVal, retOk := topStack.Pop()
	if !retOk {
		panic("invariant")
	}

	if topStack.IsEmpty() {
		stack.data.Pop()
	}
	stack.size--
	return retVal, retOk
}

func (stack *stackOfPlates[T]) Push(ele T) {
	topStack, notEmpty := stack.data.Peek()
	if notEmpty && topStack.Size() < MAX_STACK_SIZE {
		topStack.Push(ele)
	} else {
		var newStack lib.ArrayStack[T]
		newStack.Push(ele)
		stack.data.Push(&newStack)
	}
	stack.size++
}

func (stack *stackOfPlates[T]) Size() int {
	topStack, notEmpty := stack.data.Peek()

	if !notEmpty {
		return 0
	}

	return ((stack.data.Size() - 1) * MAX_STACK_SIZE) + topStack.Size()
}

// index is 0-indexed
func (stack *stackOfPlates[T]) PopAt(index int) (T, bool) {
	tmp := NewStackOfPlates[T]()

	for range stack.data.Size() - index - 1 {
		topStack, _ := stack.data.Pop()
		for !topStack.IsEmpty() {
			topEle, _ := topStack.Pop()
			tmp.Push(topEle)
		}
	}

	topStack, topStackOk := stack.data.Peek()
	if !topStackOk {
		var zero T
		return zero, false
	}

	retVal, retOk := topStack.Pop()
	if !retOk {
		var zero T
		return zero, false
	}

	if topStack.IsEmpty() {
		stack.data.Pop()
	}

	for pop, popOk := tmp.Pop(); popOk; pop, popOk = tmp.Pop() {
		stack.Push(pop)
	}

	stack.size--
	return retVal, retOk

}
