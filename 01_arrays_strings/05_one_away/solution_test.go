package oneaway

import (
	"fmt"
	"testing"
)

type TestCase struct {
	str1     string
	str2     string
	expected bool
}

func TestOneAway(t *testing.T) {

	testCases := []TestCase{
		{"", "", true}, {"", " ", true},
		{"", "  ", false}, {"a", "a", true},
		{"a", "ba", true}, {"ab", "a", true},
		{"a", "b", true}, {"ab", "ab", true},
		{"cab", "ab", true}, {"ab", "acb", true},
		{"abc", "ab", true}, {"ab", "cab", true},
		{"cab", "ca", true}, {"cab", "cb", true},
		{"cab", "dab", true}, {"cdb", "cab", true},
		{"cab", "cad", true}, {"cab", "dad", false},
		{"pale", "ple", true}, {"pales", "pale", true},
		{"pale", "bale", true}, {"pale", "bake", false},
	}

	for i, tc := range testCases {
		fmt.Printf(
			"Test %d: one_away(\"%s\", \"%s\") = ", i+1, tc.str1, tc.str2)
		result := OneAway(tc.str1, tc.str2)
		fmt.Printf("%t ", result)
		if result != tc.expected {
			fmt.Printf("expected %t, failed", tc.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
