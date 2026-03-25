package testutils

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func StackFunctionality(t *testing.T, stack lib.Stack[int]) {
	testSize := func(expected int) func(t *testing.T) {
		return func(t *testing.T) {
			t.Run("correct Size() return", func(t *testing.T) {
				got := stack.Size()
				if expected != got {
					t.Fatalf("expected %d, got %d", expected, got)
				}
			})
		}
	}

	testEmpty := func(t *testing.T) {
		t.Run("correct IsEmpty() report", func(t *testing.T) {
			expected := true
			isEmpty := stack.IsEmpty()
			if expected != isEmpty {
				t.Fatalf("expected %t, got %t", expected, isEmpty)
			}
		})

		t.Run("correct emptiness report by Peek()", func(t *testing.T) {
			expected := false
			_, ok := stack.Peek()
			if expected != ok {
				t.Fatalf("expected %t, got %t", expected, ok)
			}
		})

		t.Run("correct emptiness report by Pop()", func(t *testing.T) {
			expected := false
			_, ok := stack.Pop()
			if expected != ok {
				t.Fatalf("expected %t, got %t", expected, ok)
			}
		})

		t.Run("size is 0", testSize(0))
	}

	testPeek := func(expected int) func(t *testing.T) {
		return func(t *testing.T) {
			t.Run("correct IsEmpty() report", func(t *testing.T) {
				expected := false
				isEmpty := stack.IsEmpty()
				if expected != isEmpty {
					t.Fatalf("expected %t, got %t", expected, isEmpty)
				}
			})

			t.Run("correct Peek() return", func(t *testing.T) {
				peeked, ok := stack.Peek()
				t.Run("correct emptiness report", func(t *testing.T) {
					expected := true
					if ok != expected {
						t.Fatalf("expected %t, got %t", expected, ok)
					}
				})
				t.Run("correct element peeked", func(t *testing.T) {
					if peeked != expected {
						t.Fatalf("expected %d, got %d", expected, peeked)
					}
				})
			})

		}
	}

	testPop := func(expected int) func(t *testing.T) {
		return func(t *testing.T) {
			t.Run("correct IsEmpty() report", func(t *testing.T) {
				expected := false
				isEmpty := stack.IsEmpty()
				if expected != isEmpty {
					t.Fatalf("expected %t, got %t", expected, isEmpty)
				}
			})

			t.Run("correct Pop() return", func(t *testing.T) {
				popped, ok := stack.Pop()
				t.Run("correct emptiness report", func(t *testing.T) {
					expected := true
					if ok != expected {
						t.Fatalf("expected %t, got %t", expected, ok)
					}
				})
				t.Run("correct element popped", func(t *testing.T) {
					if popped != expected {
						t.Fatalf("expected %d, got %d", expected, popped)
					}
				})
			})
		}
	}

	t.Run("empty", testEmpty)
	stack.Push(1)
	t.Run("has first element", testPeek(1))
	t.Run("size is 1", testSize(1))
	stack.Push(2)
	t.Run("has second element", testPeek(2))
	t.Run("size is 2", testSize(2))
	t.Run("pops second element", testPop(2))
	t.Run("size is 1", testSize(1))
	t.Run("pops first element", testPop(1))
	t.Run("empty again", testEmpty)
}
