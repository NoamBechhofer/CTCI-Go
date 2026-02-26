package palindromepermutation

func init_bit_vec() uint32 {
	return 0
}

func toggle_bit(vec uint32, idx int) uint32 {
	return vec ^ (1 << idx)
}

func toggle_corresponding_bit(vec uint32, c rune) uint32 {
	if c >= 'A' && c <= 'Z' {
		return toggle_bit(vec, int(c-'A'))
	}
	if c >= 'a' && c <= 'z' {
		return toggle_bit(vec, int(c-'a'))
	}
	return vec
}

func vec_has_at_most_one_bit_set(vec uint32) bool {
	return (vec & (vec - 1)) == 0
}

func PalindromePermutation(str string) bool {
	bit_vector := init_bit_vec()
	for _, r := range str {
		bit_vector = toggle_corresponding_bit(bit_vector, r)
	}
	return vec_has_at_most_one_bit_set(bit_vector)
}
