package palindromepermutation

func bitToggled(vec uint32, idx int) uint32 {
	return vec ^ (1 << idx)
}

func correspondingBitToggled(vec uint32, c rune) uint32 {
	if c >= 'A' && c <= 'Z' {
		return bitToggled(vec, int(c-'A'))
	}
	if c >= 'a' && c <= 'z' {
		return bitToggled(vec, int(c-'a'))
	}
	return vec
}

func atMostOneBitSet(vec uint32) bool {
	return (vec & (vec - 1)) == 0
}

func PalindromePermutation(str string) bool {
	bitVector := uint32(0)
	for _, r := range str {
		bitVector = correspondingBitToggled(bitVector, r)
	}
	return atMostOneBitSet(bitVector)
}
