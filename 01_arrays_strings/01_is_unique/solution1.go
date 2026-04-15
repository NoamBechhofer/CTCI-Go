package isunique

func IsUnique(str string) bool {
	seen := map[rune]struct{}{}

	for _, r := range str {
		if _, exists := seen[r]; exists {
			return false
		}
		seen[r] = struct{}{}
	}

	return true
}
