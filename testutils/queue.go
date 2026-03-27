package testutils

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func testQueueIsEmpty[T any](t *testing.T, queue lib.Queue[T], expected bool) {
	t.Run("IsEmpty()", func(t *testing.T) {
		isEmpty := queue.IsEmpty()
		if expected != isEmpty {
			t.Fatalf("expected %t, got %t", expected, isEmpty)
		}
	})
}

func testQueuePeek[T comparable](t *testing.T, queue lib.Queue[T], expectedVal T, expectedOk bool) {
	t.Run("Peek()", func(t *testing.T) {
		peeked, ok := queue.Peek()
		t.Run("correct emptiness report", func(t *testing.T) {
			if ok != expectedOk {
				t.Fatalf("expected %t, got %t", expectedOk, ok)
			}
		})
		t.Run("correct element peeked", func(t *testing.T) {
			if peeked != expectedVal {
				t.Fatalf("expected %v, got %v", expectedVal, peeked)
			}
		})
	})
}

func testQueueRemove[T comparable](t *testing.T, queue lib.Queue[T], expectedVal T, expectedOk bool) {
	t.Run("Remove()", func(t *testing.T) {
		removed, ok := queue.Remove()
		t.Run("correct emptiness report", func(t *testing.T) {
			if ok != expectedOk {
				t.Fatalf("expected %t, got %t", expectedOk, ok)
			}
		})
		t.Run("correct element removed", func(t *testing.T) {
			if removed != expectedVal {
				t.Fatalf("expected %v, got %v", expectedVal, removed)
			}
		})
	})
}

func testQueueAdd[T any](t *testing.T, queue lib.Queue[T], val T) {
	t.Run("Add()", func(t *testing.T) {
		queue.Add(val)
	})
}

func testQueueSize[T any](t *testing.T, queue lib.Queue[T], expected int) {
	t.Run("Size()", func(t *testing.T) {
		got := queue.Size()
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func testQueueBasicOperationsAndEmpty(t *testing.T, queue lib.Queue[int]) {
	t.Run("basic queue operations test", func(t *testing.T) {
		testEmpty := func(t *testing.T) {
			testQueueIsEmpty(t, queue, true)
			testQueuePeek(t, queue, 0, false)
			testQueueRemove(t, queue, 0, false)
			testQueueSize(t, queue, 0)
		}

		testPeek := func(expected int) func(t *testing.T) {
			return func(t *testing.T) {
				testQueueIsEmpty(t, queue, false)
				testQueuePeek(t, queue, expected, true)
			}
		}

		testRemove := func(expected int) func(t *testing.T) {
			return func(t *testing.T) {
				testQueueIsEmpty(t, queue, false)
				testQueueRemove(t, queue, expected, true)
			}
		}

		t.Run("empty", testEmpty)
		queue.Add(1)
		t.Run("has first element", testPeek(1))
		testQueueSize(t, queue, 1)
		queue.Add(2)
		t.Run("has first element still", testPeek(1))
		testQueueSize(t, queue, 2)
		t.Run("removes first element", testRemove(1))
		t.Run("has second element now", testPeek(2))
		testQueueSize(t, queue, 1)
		queue.Add(3)
		t.Run("has second element still", testPeek(2))
		testQueueSize(t, queue, 2)
		t.Run("removes second element", testRemove(2))
		t.Run("has third element now", testPeek(3))
		testQueueSize(t, queue, 1)
		t.Run("removes third element", testRemove(3))
		t.Run("empty again", testEmpty)
	})
}

func testQueueAdd10kAndEmpty(t *testing.T, queue lib.Queue[int]) {
	t.Run("10k adds and 10k removes", func(t *testing.T) {
		for i := range 10_000 {
			queue.Add(i)
		}

		for i := range 10_000 {
			removed, ok := queue.Remove()
			if !ok {
				t.Fatalf("expected to remove %d, but queue was empty", i)
			}
			if removed != i {
				t.Fatalf("expected to remove %d, but removed %d", i, removed)
			}
		}
		testQueueRemove(t, queue, 0, false)
	})
}

func QueueFunctionality(t *testing.T, queue lib.Queue[int]) {
	testQueueBasicOperationsAndEmpty(t, queue)
	testQueueAdd10kAndEmpty(t, queue)
}
