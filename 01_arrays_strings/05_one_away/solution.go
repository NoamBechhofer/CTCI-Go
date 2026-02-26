package oneaway

func check_added(shorter string, longer string) bool {
	i, j := 0, 0
	found_mismatch := false
	for i < len(shorter) {
		if shorter[i] != longer[j] {
			if found_mismatch {
				return false
			}
			found_mismatch = true
			j++
			if longer[j] != shorter[i] {
				return false
			}
		}
		i++
		j++
	}

	return true
}

func check_replaced(str1 string, str2 string) bool {
	found_replacement := false

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if found_replacement {
				return false
			}
			found_replacement = true
		}
	}

	return true
}

func OneAway(str1 string, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	size_diff := len1 - len2
	switch size_diff {
	case -1:
		return check_added(str1, str2)
	case 0:
		return check_replaced(str1, str2)
	case 1:
		return check_added(str2, str1)
	default:
		return false
	}
}
