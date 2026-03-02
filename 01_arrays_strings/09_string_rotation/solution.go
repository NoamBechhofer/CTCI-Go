package stringrotation

import "strings"

func IsSubstring(super string, sub string) bool {
	return strings.Contains(super, sub)
}

func StringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return IsSubstring(strings.Repeat(s1, 2), s2)
}
