package threeinone

import (
	"sync"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

// #region struct

type StackId int

const (
	Stack1 StackId = 1
	Stack2 StackId = 2
	Stack3 StackId = 3
)

type ThreeInOne[T any] struct {
	mu    sync.Mutex
	slice []T
	// 1 pushes to right1, 2 pushes to right2, 3 pushes to left3. left2 marks
	// the beginning of stack 2. So right1, right2, and left3 point to
	// unnocupied spaces, while left2 is occupied unless stack2 is empty (in
	// which case left2 == right2)
	//
	// equivalently:
	// stack 1 is [0, right1), stack2 is [left2, right2), stack3 is reverse in (left3, len)
	right1, left2, right2, left3 int
}

// #endregion

// #region internals

func (stacks *ThreeInOne[T]) ensureInit() {
	if len(stacks.slice) == 0 {
		stacks.slice = make([]T, 9)
		stacks.right1 = 0
		stacks.left2 = 3
		stacks.right2 = 3
		stacks.left3 = 8
	}
}

func (stacks *ThreeInOne[T]) size1() int {
	return stacks.right1
}

func (stacks *ThreeInOne[T]) size2() int {
	return stacks.right2 - stacks.left2
}

func (stacks *ThreeInOne[T]) size3() int {
	return len(stacks.slice) - 1 - stacks.left3
}

func (stacks *ThreeInOne[T]) grow() {
	oldCap := len(stacks.slice)
	newCap := oldCap * 2
	newSlice := make([]T, newCap)

	newRight1 := stacks.right1
	for i := range stacks.size1() {
		newSlice[i] = stacks.slice[i]
	}

	newLeft2 := max(newRight1+1, newCap/3)
	newRight2 := newLeft2 + stacks.size2()
	for i := range stacks.size2() {
		newSlice[newLeft2+i] = stacks.slice[stacks.left2+i]
	}

	newLeft3 := stacks.left3 + (newCap - oldCap)
	for i := range stacks.size3() {
		newSlice[newLeft3+1+i] = stacks.slice[stacks.left3+1+i]
	}

	stacks.slice = newSlice
	stacks.right1 = newRight1
	stacks.left2 = newLeft2
	stacks.right2 = newRight2
	stacks.left3 = newLeft3
}

// if only one slot is available, should stack 2 move rightward (true) giving
// the slot to stack 1, or leftward (false) giving the slot to stacks 2 and 3
type shuffleDirection bool

const (
	left  = false
	right = true
)

func (stacks *ThreeInOne[T]) tryShuffle2OrGrow(direction shuffleDirection) {
	availableSlots := len(stacks.slice) - stacks.size1() - stacks.size2() - stacks.size3()
	switch availableSlots {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		stacks.grow()
	case 3:
		if direction == right {
			stacks.shuffle2Rightward()
		} else {
			stacks.shuffle2Leftward()
		}
	default:
		stacks.shuffle2ToMiddle(direction)
	}
}

func (stacks *ThreeInOne[T]) shuffle2Rightward() {
	var zero T

	for i := range stacks.size2() {
		stacks.slice[stacks.right2-i] = stacks.slice[stacks.right2-1-i]
	}
	stacks.slice[stacks.left2] = zero

	stacks.left2++
	stacks.right2++
}

func (stacks *ThreeInOne[T]) shuffle2Leftward() {
	var zero T

	for i := range stacks.size2() {
		stacks.slice[stacks.left2-1+i] = stacks.slice[stacks.left2+i]
	}
	stacks.slice[stacks.right2-1] = zero

	stacks.left2--
	stacks.right2--
}

func (stacks *ThreeInOne[T]) shuffle2ToMiddle(direction shuffleDirection) {
	availableSlots := len(stacks.slice) - stacks.size1() - stacks.size2() - stacks.size3()

	// baseline: 1/3 of free room for stack 1, 2/3 shared by stacks 2 and 3
	stack1FreeSlots := availableSlots / 3

	// but always leave room for the push that triggered the shuffle
	if direction == right && stack1FreeSlots == 0 {
		stack1FreeSlots = 1
	}

	newLeft2 := stacks.right1 + stack1FreeSlots
	newRight2 := newLeft2 + stacks.size2()

	var zero T

	if newLeft2 < stacks.left2 {
		for i := range stacks.size2() {
			stacks.slice[newLeft2+i] = stacks.slice[stacks.left2+i]
			stacks.slice[stacks.left2+i] = zero
		}
	} else if newLeft2 > stacks.left2 {
		for i := range stacks.size2() {
			stacks.slice[newRight2-1-i] = stacks.slice[stacks.right2-1-i]
			stacks.slice[stacks.right2-1-i] = zero
		}
	} else {
		panic("shuffle2ToMiddle called as no-op")
	}

	stacks.left2 = newLeft2
	stacks.right2 = newRight2
}

func createEmptinessReturn[T any]() (T, bool) {
	var zero T
	return zero, false
}

// #endregion

// #region constructor
func NewThreeInOne[T any]() *ThreeInOne[T] {
	var threeStacks ThreeInOne[T]
	threeStacks.ensureInit()
	return &threeStacks
}

// #endregion

// #region IsEmpty()
func (stacks *ThreeInOne[T]) IsEmpty(stack StackId) bool {
	stacks.mu.Lock()
	defer stacks.mu.Unlock()

	stacks.ensureInit()
	switch stack {
	case Stack1:
		return stacks.isEmpty1()
	case Stack2:
		return stacks.isEmpty2()
	case Stack3:
		return stacks.isEmpty3()
	default:
		panic("invalid stack " + lib.SignedToString(stack))
	}
}

func (stacks *ThreeInOne[T]) isEmpty1() bool {
	return stacks.right1 == 0
}
func (stacks *ThreeInOne[T]) isEmpty2() bool {
	return stacks.left2 == stacks.right2
}
func (stacks *ThreeInOne[T]) isEmpty3() bool {
	return stacks.left3 == len(stacks.slice)-1
}

// #endregion
// #region Peek()
func (stacks *ThreeInOne[T]) Peek(stack StackId) (T, bool) {
	stacks.mu.Lock()
	defer stacks.mu.Unlock()

	stacks.ensureInit()
	switch stack {
	case Stack1:
		return stacks.peek1()
	case Stack2:
		return stacks.peek2()
	case Stack3:
		return stacks.peek3()
	default:
		panic("invalid stack " + lib.SignedToString(stack))
	}
}

func (stacks *ThreeInOne[T]) peek1() (T, bool) {
	if stacks.isEmpty1() {
		return createEmptinessReturn[T]()
	}

	return stacks.slice[stacks.right1-1], true
}

func (stacks *ThreeInOne[T]) peek2() (T, bool) {
	if stacks.isEmpty2() {
		return createEmptinessReturn[T]()
	}

	return stacks.slice[stacks.right2-1], true
}

func (stacks *ThreeInOne[T]) peek3() (T, bool) {
	if stacks.isEmpty3() {
		return createEmptinessReturn[T]()
	}

	return stacks.slice[stacks.left3+1], true
}

// #endregion
// #region Pop()
func (stacks *ThreeInOne[T]) Pop(stack StackId) (T, bool) {
	stacks.mu.Lock()
	defer stacks.mu.Unlock()

	stacks.ensureInit()
	switch stack {
	case Stack1:
		return stacks.pop1()
	case Stack2:
		return stacks.pop2()
	case Stack3:
		return stacks.pop3()
	default:
		panic("invalid stack " + lib.SignedToString(stack))
	}
}

func (stacks *ThreeInOne[T]) pop1() (T, bool) {
	if stacks.isEmpty1() {
		return createEmptinessReturn[T]()
	}

	var zero T
	ret := stacks.slice[stacks.right1-1]
	stacks.slice[stacks.right1-1] = zero
	stacks.right1--
	return ret, true
}

func (stacks *ThreeInOne[T]) pop2() (T, bool) {
	if stacks.isEmpty2() {
		return createEmptinessReturn[T]()
	}

	var zero T
	ret := stacks.slice[stacks.right2-1]
	stacks.slice[stacks.right2-1] = zero
	stacks.right2--
	return ret, true
}

func (stacks *ThreeInOne[T]) pop3() (T, bool) {
	if stacks.isEmpty3() {
		return createEmptinessReturn[T]()
	}

	var zero T
	ret := stacks.slice[stacks.left3+1]
	stacks.slice[stacks.left3+1] = zero
	stacks.left3++
	return ret, true
}

// #endregion
// #region Push()
func (stacks *ThreeInOne[T]) Push(stack StackId, ele T) {
	stacks.mu.Lock()
	defer stacks.mu.Unlock()

	stacks.ensureInit()
	switch stack {
	case Stack1:
		stacks.push1(ele)
	case Stack2:
		stacks.push2(ele)
	case Stack3:
		stacks.push3(ele)
	default:
		panic("invalid stack " + lib.SignedToString(stack))
	}
}

func (stacks *ThreeInOne[T]) push1(ele T) {
	if stacks.right1 >= stacks.left2 {
		stacks.tryShuffle2OrGrow(right)
	}

	stacks.slice[stacks.right1] = ele
	stacks.right1++
}

func (stacks *ThreeInOne[T]) push2(ele T) {
	if stacks.right2 >= stacks.left3+1 {
		stacks.tryShuffle2OrGrow(left)
	}

	stacks.slice[stacks.right2] = ele
	stacks.right2++
}

func (stacks *ThreeInOne[T]) push3(ele T) {
	if stacks.left3 <= stacks.right2-1 {
		stacks.tryShuffle2OrGrow(left)
	}

	stacks.slice[stacks.left3] = ele
	stacks.left3--
}

// #endregion

// #region Size()
func (stacks *ThreeInOne[T]) Size(stack StackId) int {
	stacks.mu.Lock()
	defer stacks.mu.Unlock()

	stacks.ensureInit()
	switch stack {
	case Stack1:
		return stacks.size1()
	case Stack2:
		return stacks.size2()
	case Stack3:
		return stacks.size3()
	default:
		panic("invalid stack " + lib.SignedToString(stack))
	}
}

// #endregion
