package isunique

import (
	"fmt"
	"testing"
)

type TestCase struct {
	str      string
	expected bool
}

func TestIsUnique(t *testing.T) {
	testCases := [...]TestCase{
		{"", true},
		{"a", true},
		{"abcde", true},
		{"hello", false},
		{"1234567890", true},
		{"AaBbCc", true},
		{"AaBbCcA", false},
		{"!@#$%^&*()", true},
		{"!@#$%^&*()!", false},
		{"😊😍🤷‍♂️😉", true},
	}

	for _, testCase := range testCases {
		result := IsUnique(testCase.str)
		fmt.Printf("IsUnique(%s) = %t, ", testCase.str, result)
		if !result == testCase.expected {
			fmt.Printf("expected %t, failed\n", testCase.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}

	for _, testCase := range testCases {
		result := IsUniqueNoAdditionalDataStructures([]rune(testCase.str))
		fmt.Printf("IsUniqueNoAdditionalDataStructures(%s) = %t, ", testCase.str, result)
		if !result == testCase.expected {
			fmt.Printf("expected %t, failed\n", testCase.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
