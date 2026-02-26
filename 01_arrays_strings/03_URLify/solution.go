package urlify

func URLify(str []rune, trueLength int) {
	right := len(str) - 1

	for left := trueLength - 1; left >= 0; left-- {
		switch str[left] {
		case ' ':
			str[right] = '0'
			right--
			str[right] = '2'
			right--
			str[right] = '%'
			right--
		default:
			str[right] = str[left]
			right--
		}
	}
}
