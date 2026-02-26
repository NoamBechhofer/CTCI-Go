package isunique

import "math"

const HASH_TABLE_INIT_SIZE = 11
const HASH_TABLE_MAX_LOAD_FACTOR = 0.7

// isPrime reports whether a number that is at least HASH_TABLE_INIT_SIZE is
// prime.
// This is not a generalized isPrime, as it does not work for small numbers.
func isPrime(n int) bool {
	// save CPU instructions by ignoring n == 1 or 2 which will never be called
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func selectPrime(n int) int {
	for i := n*2 + 1; ; i += 2 {
		if isPrime(i) {
			return i
		}
	}
}

func hash(r rune) int { return int(r) }

type HashTable struct {
	size  int
	table []rune
}

// specialized HashTable for IsUnique
func NewHashTable() *HashTable {
	ht := new(HashTable)
	ht.make(HASH_TABLE_INIT_SIZE)
	return ht
}

func (ht *HashTable) make(capacity int) {
	ht.size = 0
	ht.table = make([]rune, capacity)
	for i := range ht.table {
		ht.table[i] = -1
	}
}

func (ht *HashTable) loadFactor() float32 {
	return float32(ht.size) / float32(len(ht.table))
}

func (ht *HashTable) probe(r rune, iteration int) int {
	// use linear probing for cache locality
	return (hash(r) + iteration) % len(ht.table)
}

func (ht *HashTable) rehash() {
	oldTable := ht.table

	ht.make(selectPrime(len(ht.table)))

	for i := range oldTable {
		r := oldTable[i]
		if r != -1 {
			for j := 0; j < len(ht.table); j++ {
				slot := ht.probe(r, j)
				if ht.table[slot] == -1 {
					ht.table[slot] = r
					ht.size++
					break
				}
			}
		}
	}
}

func (ht *HashTable) IsOk(r rune) bool {
	for i := 0; i < len(ht.table); i++ {
		slot := ht.probe(r, i)
		switch ht.table[slot] {
		case -1:
			ht.table[slot] = r
			ht.size++
			if ht.loadFactor() > HASH_TABLE_MAX_LOAD_FACTOR {
				ht.rehash()
			}
			return true
		case r:
			return false
		default:
			continue
		}
	}
	panic("should return by now")
}

// IsUnique reports whether all characters in s are unique.
func IsUnique(str string) bool {
	seen := NewHashTable()

	for _, r := range str {
		if !seen.IsOk(r) {
			return false
		}
	}

	return true
}
