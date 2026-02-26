package urlify

import (
	"fmt"
	"slices"
	"testing"
)

type TestCase struct {
	input       []rune
	trueLength int
	expected    []rune
}

func TestURLify(t *testing.T) {

	testCases := []TestCase{
		{[]rune(""), 0, []rune("")},
		{[]rune("a"), 1, []rune("a")},
		{[]rune("   "), 1, []rune("%20")},
		{[]rune("Mr John Smith    "), 13, []rune("Mr%20John%20Smith")},
		{[]rune("   leading spaces        "), 17, []rune("%20%20%20leading%20spaces")},
		{[]rune("trailing spaces           "), 18, []rune("trailing%20spaces%20%20%20")},
	}

	for i, tc := range testCases {
		fmt.Printf("Test %d: urlify(%q, %d) = ",
			i+1,
			tc.input,
			tc.trueLength,
		)
		URLify(tc.input, tc.trueLength)
		fmt.Printf("%q ", tc.input)
		if !slices.Equal(tc.input, tc.expected) {
			fmt.Printf("expected %q, failed\n", tc.expected)
			t.Fail()
		} else {
			fmt.Printf("passed\n")
		}
	}
}
