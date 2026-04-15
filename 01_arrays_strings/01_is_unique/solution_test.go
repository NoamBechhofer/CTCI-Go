package isunique

import (
	"fmt"
	"testing"
)

type testCase struct {
	str  string
	want bool
}

func TestIsUnique(t *testing.T) {
	testCases := []testCase{
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

	solutions := []struct {
		f    func(str string) bool
		name string
	}{
		{name: "IsUnique", f: IsUnique},
		{name: "IsUniqueNoAdditionalDataStructures", f: func(str string) bool { return IsUniqueNoAdditionalDataStructures([]rune(str)) }},
	}

	for _, solution := range solutions {
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%s(%s)", solution.name, testCase.str), func(t *testing.T) {
				got := solution.f(testCase.str)
				if got != testCase.want {
					t.Fatalf("expected %t, got %t", testCase.want, got)
				} else {
					t.Logf("got %t", got)
				}
			})
		}
	}
}
