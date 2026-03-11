package lib

import (
	"container/list"
	"fmt"
)

type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

type SinglyLinkedList[T any] struct {
	Head *SinglyLinkedListNode[T]
}

func SinglyLinkedListFromSlice[T any](slice []T) SinglyLinkedList[T] {
	if len(slice) == 0 {
		return SinglyLinkedList[T]{Head: nil}
	}

	block := make([]SinglyLinkedListNode[T], len(slice))

	for i := range len(block) - 1 {
		block[i] = SinglyLinkedListNode[T]{Val: slice[i], Next: &block[i+1]}
	}
	block[len(block)-1] = SinglyLinkedListNode[T]{Val: slice[len(slice)-1], Next: nil}

	return SinglyLinkedList[T]{Head: &block[0]}
}

func (list *SinglyLinkedList[T]) ToSlice() []T {
	ret := []T{}

	for curr := list.Head; curr != nil; curr = curr.Next {
		ret = append(ret, curr.Val)
	}

	return ret
}

type DoublyLinkedListNode[T any] struct {
	Val  T
	Next *DoublyLinkedListNode[T]
	Prev *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T any] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
}

func DoublyLinkedListFromSlice[T any](slice []T) DoublyLinkedList[T] {
	if len(slice) == 0 {
		return DoublyLinkedList[T]{Head: nil, Tail: nil}
	}

	if len(slice) == 1 {
		node := DoublyLinkedListNode[T]{Prev: nil, Val: slice[0], Next: nil}
		return DoublyLinkedList[T]{Head: &node, Tail: &node}
	}

	block := make([]DoublyLinkedListNode[T], len(slice))

	block[0] = DoublyLinkedListNode[T]{Prev: nil, Val: slice[0], Next: &block[1]}

	for i := 1; i < len(block)-1; i++ {
		block[i] = DoublyLinkedListNode[T]{Prev: &block[i-1], Val: slice[i], Next: &block[i+1]}
	}

	block[len(block)-1] = DoublyLinkedListNode[T]{Prev: &block[len(block)-2], Val: slice[len(slice)-1], Next: nil}

	return DoublyLinkedList[T]{Head: &block[0], Tail: &block[len(block)-1]}
}

func (list *DoublyLinkedList[T]) ToSlice() []T {
	ret := []T{}

	for curr := list.Head; curr != nil; curr = curr.Next {
		ret = append(ret, curr.Val)
	}

	return ret
}

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
