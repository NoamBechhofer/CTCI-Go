package oneaway

import (
	"fmt"
	"testing"
)

type testCase struct {
	str1 string
	str2 string
	want bool
}

func TestOneAway(t *testing.T) {
	testCases := []testCase{
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

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("\"%s\", \"%s\"", tc.str1, tc.str2), func(t *testing.T) {
			got := OneAway(tc.str1, tc.str2)
			if got != tc.want {
				t.Fatalf("wanted %t, got %t", tc.want, got)
			} else {
				t.Logf("%t ", got)
			}
		})
	}
}
