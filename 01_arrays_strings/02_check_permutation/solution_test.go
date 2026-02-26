package checkpermutation

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s1       string
	s2       string
	expected bool
}

func TestCheckPermutation(t *testing.T) {

	test_cases := []TestCase{{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "bac", true},
		{"aabbc", "bcaba", true},
		{"abc", "ab", false},
		{"abcde", "edcba", true},
		{"hello", "oellh", true},
		{"1234567890", "0987654321", true},
		{"AaBbCc", "bBcCaA", true},
		{"AaBbCcA", "aAbBcCA", true},
		{"!@#$%^&*()", ")(*&^%$#@!", true},
		{"!@#$%^&*()!", "!@#$%^&*()!", true},
		{"abcde", "abcd", false},
		{"hello", "helloo", false},
		{"1234567890", "123456890", false},
		{"AaBbCc", "AaBbCcc", false},
		{"!@#$%^&*()", "!@#$%^&*() ", false},
		{"😹emojicats😺😻😸😼😽🙀", "😽😺tim😸oa🙀sj😹😼😻ce", true},
		{"😹emojicats😺😻😸😼😽🙀", "😽😽tim😸oa🙀sj😹😼😻ce", false}}

	for _, tc := range test_cases {
		result := CheckPermutation(tc.s1, tc.s2)
		fmt.Printf("CheckPermutation(%s, %s) = %t, ", tc.s1, tc.s2, result)
		if !result == tc.expected {
			fmt.Printf("expected %t, failed\n", tc.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
