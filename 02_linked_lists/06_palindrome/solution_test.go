package palindrome

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	str  string
	want bool
}

func TestPalindrome(t *testing.T) {

	testCases := []TestCase{
		// empty / minimal
		{str: "", want: true},
		{str: "a", want: true},

		// length 2
		{str: "aa", want: true},
		{str: "ab", want: false},

		// length 3
		{str: "aaa", want: true},
		{str: "aab", want: false},
		{str: "aba", want: true},
		{str: "abb", want: false},
		{str: "abc", want: false},

		// length 4
		{str: "abba", want: true},
		{str: "aaaa", want: true},
		{str: "abca", want: false},
		{str: "abab", want: false},
		{str: "abcc", want: false},

		// length 5
		{str: "abcba", want: true},
		{str: "aaaaa", want: true},
		{str: "ababa", want: true},
		{str: "abcaa", want: false},
		{str: "abcca", want: false},

		// longer palindromes
		{str: "racecar", want: true},
		{str: "level", want: true},
		{str: "rotator", want: true},
		{str: "madam", want: true},

		// longer non-palindromes
		{str: "palindrome", want: false},
		{str: "abcdefg", want: false},
		{str: "racecars", want: false},

		// repeated patterns
		{str: "aaaaaa", want: true},
		{str: "aaaaba", want: false},
		{str: "baaaab", want: true},

		// mixed characters
		{str: "a1a", want: true},
		{str: "1a1", want: true},
		{str: "a1b", want: false},
		{str: "1221", want: true},
		{str: "1231", want: false},
	}

	solutions := []struct {
		name string
		f    func(list lib.SinglyLinkedList[rune]) bool
	}{
		{name: "Palindrome", f: Palindrome[rune]},
		{name: "PalindromeRecursive", f: PalindromeRecursive[rune]},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			testName := fmt.Sprintf("%s(%s)", solution.name, string(tc.str))
			testFunc := func(t *testing.T) {
				list := lib.SinglyLinkedListFromSlice([]rune(tc.str))
				got := solution.f(list)
				if got != tc.want {
					t.Fatalf("want %t, got %t", tc.want, got)
				} else {
					t.Logf("got %t", got)
				}
			}
			t.Run(testName, testFunc)
		}
	}
}
