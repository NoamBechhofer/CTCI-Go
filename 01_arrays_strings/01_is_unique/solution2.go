package isunique

import "math/rand/v2"

// partition runs a hoare partition, partitioning runes according to the
// folowing rules:
//
// - pivot is chosen as the first element before partition.
//
// - after partition an index p is returned such that all elements inclusively
// left of p are <= the pivot, and all elements exlusively right of p are >=
// the pivot
func partition(runes []rune) int {
	pivot := runes[0]

	left := -1
	right := len(runes)

	for {
		for {
			left++
			if runes[left] >= pivot {
				break
			}
		}
		for {
			right--
			if runes[right] <= pivot {
				break
			}
		}
		if left >= right {
			return right
		}
		runes[left], runes[right] = runes[right], runes[left]
	}
}

func quickSort(runes []rune) {
	switch len(runes) {
	case 0:
		fallthrough
	case 1:
		return
	case 2:
		if runes[0] > runes[1] {
			runes[0], runes[1] = runes[1], runes[0]
		}
		return
	}

	pivotIdx := rand.IntN(len(runes))
	runes[0], runes[pivotIdx] = runes[pivotIdx], runes[0]

	partitionIdx := partition(runes)
	quickSort(runes[:partitionIdx+1])
	quickSort(runes[partitionIdx+1:])
}

func inPlaceSort(runes []rune) {
	quickSort(runes)
}

// IsUniqueNoAdditionalDataStructures determines if s has all unique characters.
// Reports whether all characters in s are unique. No additional data structures
// are used, so the parameter cannot be a string (strings are immutable in go)
func IsUniqueNoAdditionalDataStructures(runes []rune) bool {
	if len(runes) < 2 {
		return true
	}

	inPlaceSort(runes)

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}

	return true
}
