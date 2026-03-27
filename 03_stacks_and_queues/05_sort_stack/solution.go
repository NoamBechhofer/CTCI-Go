package sortstack

import (
	"cmp"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type constrainedStack[T any] struct {
	data lib.Stack[T]
}

func (stack *constrainedStack[T]) IsEmpty() bool   { return stack.data.IsEmpty() }
func (stack *constrainedStack[T]) Peek() (T, bool) { return stack.data.Peek() }
func (stack *constrainedStack[T]) Pop() (T, bool)  { return stack.data.Pop() }
func (stack *constrainedStack[T]) Push(ele T)      { stack.data.Push(ele) }

func SortStack[T cmp.Ordered](stack lib.Stack[T]) {
	var tmpStack lib.ArrayStack[T]

	wrapper := constrainedStack[T]{data: stack}
	tmp := constrainedStack[T]{data: &tmpStack}

	for pop, ok := stack.Pop(); ok; pop, ok = stack.Pop() {
		for {
			if tmp.IsEmpty() {
				tmp.Push(pop)
				break
			}

			tmpPeek, _ := tmp.Peek()
			if tmpPeek <= pop {
				tmp.Push(pop)
				break
			}

			tmpPop, _ := tmp.Pop()
			wrapper.Push(tmpPop)

		}
	}

	for !tmp.IsEmpty() {
		pop, _ := tmp.Pop()
		wrapper.Push(pop)
	}
}
