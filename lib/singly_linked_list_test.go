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
