package oneaway

func checkAdded(shorter string, longer string) bool {
	foundMismatch := false
	for i, j := 0, 0; i < len(shorter); i, j = i+1, j+1 {
		if shorter[i] != longer[j] {
			if foundMismatch {
				return false
			}
			foundMismatch = true
			j++
			if longer[j] != shorter[i] {
				return false
			}
		}
	}

	return true
}

func checkReplaced(str1 string, str2 string) bool {
	foundReplacement := false
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if foundReplacement {
				return false
			}
			foundReplacement = true
		}
	}

	return true
}

func OneAway(str1 string, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	sizeDiff := len1 - len2
	switch sizeDiff {
	case -1:
		return checkAdded(str1, str2)
	case 0:
		return checkReplaced(str1, str2)
	case 1:
		return checkAdded(str2, str1)
	default:
		return false
	}
}
