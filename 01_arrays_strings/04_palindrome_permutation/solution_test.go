package palindromepermutation

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected bool
}

func TestPalindromePermutation(t *testing.T) {

	testCases := []TestCase{{"", true},
		{" ", true},
		{"a ", true},
		{" a", true},
		{"a", true},
		{"ab ", false},
		{" ab", false},
		{"ab", false},
		{"ba ", false},
		{" ba", false},
		{"ba", false},
		{"baa", true},
		{"aba", true},
		{"aab", true},
		{" baa", true},
		{" aba", true},
		{" aab", true},
		{"b aa", true},
		{"a ba", true},
		{"a ab", true},
		{"ba a", true},
		{"ab a", true},
		{"aa b", true},
		{"baa ", true},
		{"aba ", true},
		{"aab ", true},
		{"abc", false},
		{" a a b b  ", true},
		{" a a b b c ", true},
		{" a a d b b c ", false}}

	for i, tc := range testCases {
		fmt.Printf("Test %d: palindrome_permutation(\"%s\") = ", i+1, tc.input)
		result := PalindromePermutation(tc.input)
		fmt.Printf("%t ", result)
		if !result == tc.expected {
			fmt.Printf("expected \"%t\", failed", tc.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
