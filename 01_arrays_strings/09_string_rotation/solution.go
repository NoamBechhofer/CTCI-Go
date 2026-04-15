package stringrotation

import "strings"

func StringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(strings.Repeat(s1, 2), s2)
}
