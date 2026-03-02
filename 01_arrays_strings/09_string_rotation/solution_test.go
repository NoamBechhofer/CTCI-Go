package stringrotation

import (
	"fmt"
	"testing"
)

type TestCase struct {
	str1     string
	str2     string
	expected bool
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
		result := StringRotation(testCase.str1, testCase.str2)
		fmt.Printf("IsUnique(%q, %q) = %t, ", testCase.str1, testCase.str2, result)
		if !result == testCase.expected {
			fmt.Printf("expected %t, failed\n", testCase.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
