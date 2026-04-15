package stringrotation

import (
	"fmt"
	"testing"
)

type TestCase struct {
	str1 string
	str2 string
	want bool
}

func TestStringRotation(t *testing.T) {
	testCases := []TestCase{
		{"waterbottle", "erbottlewat", true},
		{"waterbottle", "erbottlewta", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"hello", "oellh", false},
		{"hello", "hello", true},
		{"hello", "llohe", true},
		{"hello", "llo", false},
		{"hello", "world", false},
		{"hello", "oehll", false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%q, %q", testCase.str1, testCase.str2), func(t *testing.T) {
			got := StringRotation(testCase.str1, testCase.str2)
			if got != testCase.want {
				t.Fatalf("wanted %t, got %t", testCase.want, got)
			} else {
				t.Logf("got %t", got)
			}
		})
	}
}
