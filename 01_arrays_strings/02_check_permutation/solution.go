package checkpermutation

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hm := map[rune]int{}

	for _, r := range s1 {
		hm[r]++
	}

	for _, r := range s2 {
		_, exists := hm[r]
		if !exists {
			return false
		}
		hm[r]--
		if hm[r] < 0 {
			return false
		}
	}
	return true
}
