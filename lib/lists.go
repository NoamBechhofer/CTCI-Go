package lib

import (
	"container/list"
	"fmt"
)

type TypedElement[T any] struct {
	element *list.Element
}

func typedElementOf[T any](element *list.Element) *TypedElement[T] {
	if element == nil {
		return nil
	}
	if _, ok := element.Value.(T); !ok {
		panic(fmt.Sprintf("tried to make TypedElement[%T] from %T", *new(T), element.Value))
	}
	return &TypedElement[T]{element: element}
}

func (e *TypedElement[T]) Value() T {
	if e == nil || e.element == nil {
		panic("nil element")
	}
	v, ok := e.element.Value.(T)
	if !ok {
		panic("stored value has unexpected type")
	}
	return v
}

func (e *TypedElement[T]) Next() *TypedElement[T] {
	if e == nil || e.element == nil {
		return nil
	}
	return typedElementOf[T](e.element.Next())
}

func (e *TypedElement[T]) Prev() *TypedElement[T] {
	if e == nil || e.element == nil {
		return nil
	}
	return typedElementOf[T](e.element.Prev())
}

type TypedList[T any] struct {
	list *list.List
}

func NewTypedList[T any]() *TypedList[T] {
	return &TypedList[T]{list: list.New()}
}

func (l *TypedList[T]) Back() *TypedElement[T]  { return typedElementOf[T](l.list.Back()) }
func (l *TypedList[T]) Front() *TypedElement[T] { return typedElementOf[T](l.list.Front()) }
func (l *TypedList[T]) Len() int                { return l.list.Len() }

func (l *TypedList[T]) Init() *TypedList[T] {
	l.list.Init()
	return l
}

func (l *TypedList[T]) InsertAfter(v T, mark *TypedElement[T]) *TypedElement[T] {
	return typedElementOf[T](l.list.InsertAfter(v, mark.element))
}

func (l *TypedList[T]) InsertBefore(v T, mark *TypedElement[T]) *TypedElement[T] {
	return typedElementOf[T](l.list.InsertBefore(v, mark.element))
}

func (l *TypedList[T]) MoveAfter(e, mark *TypedElement[T]) { l.list.MoveAfter(e.element, mark.element) }
func (l *TypedList[T]) MoveBefore(e, mark *TypedElement[T]) {
	l.list.MoveBefore(e.element, mark.element)
}
func (l *TypedList[T]) MoveToBack(e *TypedElement[T])  { l.list.MoveToBack(e.element) }
func (l *TypedList[T]) MoveToFront(e *TypedElement[T]) { l.list.MoveToFront(e.element) }

func (l *TypedList[T]) PushBack(v T) *TypedElement[T]  { return typedElementOf[T](l.list.PushBack(v)) }
func (l *TypedList[T]) PushFront(v T) *TypedElement[T] { return typedElementOf[T](l.list.PushFront(v)) }

func (l *TypedList[T]) PushBackList(other *TypedList[T])  { l.list.PushBackList(other.list) }
func (l *TypedList[T]) PushFrontList(other *TypedList[T]) { l.list.PushFrontList(other.list) }

func (l *TypedList[T]) Remove(e *TypedElement[T]) T {
	v, ok := l.list.Remove(e.element).(T)
	if !ok {
		panic("stored value has unexpected type")
	}
	return v
}

func ListFromSlice[T any](slice []T) *TypedList[T] {
	ret := NewTypedList[T]()
	for _, ele := range slice {
		ret.PushBack(ele)
	}
	return ret
}

func ListToSlice[T any](l *TypedList[T]) []T {
	ret := make([]T, 0, l.Len())
	for curr := l.Front(); curr != nil; curr = curr.Next() {
		ret = append(ret, curr.Value())
	}
	return ret
}
