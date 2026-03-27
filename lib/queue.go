package lib

type Queue[T any] interface {
	Add(ele T)
	IsEmpty() bool
	Peek() (T, bool)
	Remove() (T, bool)
	Size() int
}
