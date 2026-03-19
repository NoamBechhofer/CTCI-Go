package palindrome

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	str      string
	expected bool
}

func TestPalindrome(t *testing.T) {

	testCases := []TestCase{
		// empty / minimal
		{str: "", expected: true},
		{str: "a", expected: true},

		// length 2
		{str: "aa", expected: true},
		{str: "ab", expected: false},

		// length 3
		{str: "aaa", expected: true},
		{str: "aab", expected: false},
		{str: "aba", expected: true},
		{str: "abb", expected: false},
		{str: "abc", expected: false},

		// length 4
		{str: "abba", expected: true},
		{str: "aaaa", expected: true},
		{str: "abca", expected: false},
		{str: "abab", expected: false},
		{str: "abcc", expected: false},

		// length 5
		{str: "abcba", expected: true},
		{str: "aaaaa", expected: true},
		{str: "ababa", expected: true},
		{str: "abcaa", expected: false},
		{str: "abcca", expected: false},

		// longer palindromes
		{str: "racecar", expected: true},
		{str: "level", expected: true},
		{str: "rotator", expected: true},
		{str: "madam", expected: true},

		// longer non-palindromes
		{str: "palindrome", expected: false},
		{str: "abcdefg", expected: false},
		{str: "racecars", expected: false},

		// repeated patterns
		{str: "aaaaaa", expected: true},
		{str: "aaaaba", expected: false},
		{str: "baaaab", expected: true},

		// mixed characters
		{str: "a1a", expected: true},
		{str: "1a1", expected: true},
		{str: "a1b", expected: false},
		{str: "1221", expected: true},
		{str: "1231", expected: false},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("Palindrome(%s)", string(tc.str))
		testFunc := func(t *testing.T) {
			list := lib.SinglyLinkedListFromSlice([]rune(tc.str))
			got := Palindrome(list)
			if got != tc.expected {
				t.Fatalf("expect %t, got %t", tc.expected, got)
			} else {
				t.Logf("got %t", got)
			}
		}
		t.Run(testName, testFunc)
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("PalindromeRecursive(%s)", string(tc.str))
		testFunc := func(t *testing.T) {
			list := lib.SinglyLinkedListFromSlice([]rune(tc.str))
			got := PalindromeRecursive(list)
			if got != tc.expected {
				t.Fatalf("expect %t, got %t", tc.expected, got)
			} else {
				t.Logf("got %t", got)
			}
		}
		t.Run(testName, testFunc)
	}
}
