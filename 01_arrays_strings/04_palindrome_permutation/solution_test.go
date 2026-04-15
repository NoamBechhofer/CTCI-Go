package palindromepermutation

import (
	"testing"
)

type TestCase struct {
	input string
	want  bool
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

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := PalindromePermutation(tc.input)
			if got != tc.want {
				t.Fatalf("want %t, got %t", tc.want, got)
			} else {
				t.Logf("%t ", got)
			}
		})
	}
}
