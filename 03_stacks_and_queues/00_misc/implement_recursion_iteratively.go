package misc

import "github.com/NoamBechhofer/CTCI-Go/lib"

func FactorialRecursive(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func FactorialIterative(n uint) uint {
	type Frame struct {
		n             uint
		beforeRecurse bool
	}

	var stack lib.ArrayStack[Frame]

	stack.Push(Frame{n: n, beforeRecurse: true})

	var ret uint = 1

	for frame, ok := stack.Pop(); ok; frame, ok = stack.Pop() {
		if frame.n == 0 || frame.n == 1 {
			ret = 1
			continue
		}

		if frame.beforeRecurse {
			stack.Push(Frame{n: frame.n, beforeRecurse: false})
			stack.Push(Frame{n: frame.n - 1, beforeRecurse: true})
		} else {
			ret *= frame.n
		}
	}

	return ret
}
