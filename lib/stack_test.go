package lib

import "testing"

func TestStack(t *testing.T) {
	var stack Stack[int]

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
	stack.Push(2)
	t.Run("has second element", testPeek(2))
	t.Run("pops second element", testPop(2))
	t.Run("pops first element", testPop(1))
	t.Run("empty again", testEmpty)
}
