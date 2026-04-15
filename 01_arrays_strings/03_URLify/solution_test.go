package urlify

import (
	"fmt"
	"slices"
	"testing"
)

type testCase struct {
	input      []rune
	trueLength int
	want       []rune
}

func TestURLify(t *testing.T) {
	testCases := []testCase{
		{[]rune(""), 0, []rune("")},
		{[]rune("a"), 1, []rune("a")},
		{[]rune("   "), 1, []rune("%20")},
		{[]rune("Mr John Smith    "), 13, []rune("Mr%20John%20Smith")},
		{[]rune("   leading spaces        "), 17, []rune("%20%20%20leading%20spaces")},
		{[]rune("trailing spaces           "), 18, []rune("trailing%20spaces%20%20%20")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q, true length %d", tc.input, tc.trueLength), func(t *testing.T) {
			URLify(tc.input, tc.trueLength)
			if !slices.Equal(tc.input, tc.want) {
				t.Fatalf("expected %q, got %q\n", tc.want, tc.input)
			} else {
				t.Logf("got %q", tc.input)
			}
		})
	}
}
