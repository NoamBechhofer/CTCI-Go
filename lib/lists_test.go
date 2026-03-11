package lib

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedListFromSlice_Empty(t *testing.T) {
	list := SinglyLinkedListFromSlice([]int(nil))

	if list.Head != nil {
		t.Fatalf("expected nil head, got %#v", list.Head)
	}

	got := list.ToSlice()
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}

func TestSinglyLinkedListFromSlice_Single(t *testing.T) {
	input := []int{42}
	list := SinglyLinkedListFromSlice(input)

	if list.Head == nil {
		t.Fatal("expected non-nil head")
	}
	if list.Head.Val != 42 {
		t.Fatalf("expected head value 42, got %v", list.Head.Val)
	}
	if list.Head.Next != nil {
		t.Fatalf("expected nil next, got %#v", list.Head.Next)
	}

	got := list.ToSlice()
	if !reflect.DeepEqual(got, input) {
		t.Fatalf("expected %#v, got %#v", input, got)
	}
}

func TestSinglyLinkedListFromSlice_Multiple(t *testing.T) {
	input := []int{1, 2, 3, 4}
	list := SinglyLinkedListFromSlice(input)

	if list.Head == nil {
		t.Fatal("expected non-nil head")
	}

	values := []int{}
	for curr := list.Head; curr != nil; curr = curr.Next {
		values = append(values, curr.Val)
	}

	if !reflect.DeepEqual(values, input) {
		t.Fatalf("expected traversal %#v, got %#v", input, values)
	}

	got := list.ToSlice()
	if !reflect.DeepEqual(got, input) {
		t.Fatalf("expected %#v, got %#v", input, got)
	}
}

func TestSinglyLinkedListToSlice_IndependenceFromInputSlice(t *testing.T) {
	input := []int{1, 2, 3}
	list := SinglyLinkedListFromSlice(input)

	input[0] = 99
	input[1] = 88
	input[2] = 77

	got := list.ToSlice()
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected list values %#v after input mutation, got %#v", want, got)
	}
}

func TestSinglyLinkedListToSlice_NilReceiverHead(t *testing.T) {
	list := SinglyLinkedList[int]{Head: nil}

	got := list.ToSlice()
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}

func TestDoublyLinkedListFromSlice_Empty(t *testing.T) {
	list := DoublyLinkedListFromSlice([]int(nil))

	if list.Head != nil {
		t.Fatalf("expected nil head, got %#v", list.Head)
	}
	if list.Tail != nil {
		t.Fatalf("expected nil tail, got %#v", list.Tail)
	}

	got := list.ToSlice()
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}

func TestDoublyLinkedListFromSlice_Single(t *testing.T) {
	input := []int{42}
	list := DoublyLinkedListFromSlice(input)

	if list.Head == nil {
		t.Fatal("expected non-nil head")
	}
	if list.Tail == nil {
		t.Fatal("expected non-nil tail")
	}
	if list.Head != list.Tail {
		t.Fatal("expected head and tail to be the same node")
	}
	if list.Head.Val != 42 {
		t.Fatalf("expected value 42, got %v", list.Head.Val)
	}
	if list.Head.Prev != nil {
		t.Fatalf("expected nil prev, got %#v", list.Head.Prev)
	}
	if list.Head.Next != nil {
		t.Fatalf("expected nil next, got %#v", list.Head.Next)
	}

	got := list.ToSlice()
	if !reflect.DeepEqual(got, input) {
		t.Fatalf("expected %#v, got %#v", input, got)
	}
}

func TestDoublyLinkedListFromSlice_Multiple_ForwardAndBackwardLinks(t *testing.T) {
	input := []int{10, 20, 30, 40}
	list := DoublyLinkedListFromSlice(input)

	if list.Head == nil {
		t.Fatal("expected non-nil head")
	}
	if list.Tail == nil {
		t.Fatal("expected non-nil tail")
	}

	forward := []int{}
	var prev *DoublyLinkedListNode[int]
	for curr := list.Head; curr != nil; curr = curr.Next {
		if curr.Prev != prev {
			t.Fatalf("expected prev pointer %#v, got %#v", prev, curr.Prev)
		}
		forward = append(forward, curr.Val)
		prev = curr
	}

	if !reflect.DeepEqual(forward, input) {
		t.Fatalf("expected forward traversal %#v, got %#v", input, forward)
	}

	backward := []int{}
	var next *DoublyLinkedListNode[int]
	for curr := list.Tail; curr != nil; curr = curr.Prev {
		if curr.Next != next {
			t.Fatalf("expected next pointer %#v, got %#v", next, curr.Next)
		}
		backward = append(backward, curr.Val)
		next = curr
	}

	wantBackward := []int{40, 30, 20, 10}
	if !reflect.DeepEqual(backward, wantBackward) {
		t.Fatalf("expected backward traversal %#v, got %#v", wantBackward, backward)
	}

	got := list.ToSlice()
	if !reflect.DeepEqual(got, input) {
		t.Fatalf("expected %#v, got %#v", input, got)
	}
}

func TestDoublyLinkedListToSlice_IndependenceFromInputSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	list := DoublyLinkedListFromSlice(input)

	input[0] = "x"
	input[1] = "y"
	input[2] = "z"

	got := list.ToSlice()
	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected list values %#v after input mutation, got %#v", want, got)
	}
}

func TestDoublyLinkedListToSlice_NilReceiverHead(t *testing.T) {
	list := DoublyLinkedList[int]{Head: nil, Tail: nil}

	got := list.ToSlice()
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %#v", got)
	}
}

func TestLinkedLists_WithStructType(t *testing.T) {
	type item struct {
		ID   int
		Name string
	}

	input := []item{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}

	single := SinglyLinkedListFromSlice(input)
	double := DoublyLinkedListFromSlice(input)

	if got := single.ToSlice(); !reflect.DeepEqual(got, input) {
		t.Fatalf("singly: expected %#v, got %#v", input, got)
	}

	if got := double.ToSlice(); !reflect.DeepEqual(got, input) {
		t.Fatalf("doubly: expected %#v, got %#v", input, got)
	}
}
