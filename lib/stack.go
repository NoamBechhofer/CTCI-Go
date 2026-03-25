package lib

type Stack[T any] interface {
	IsEmpty() bool
	Peek() (T, bool)
	Pop() (T, bool)
	Push(ele T)
	Size() int
}
