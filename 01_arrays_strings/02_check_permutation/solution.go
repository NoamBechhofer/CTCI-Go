package checkpermutation

const HASH_MAP_INIT_SIZE = 11
const HASH_MAP_MAX_LOAD_FACTOR = 0.7

func hash(r rune) int {
	return int(r)
}

type HashMapEntry struct {
	r    rune
	freq int
}

func blankHashMapEntry() HashMapEntry     { return HashMapEntry{r: -1, freq: -1} }
func (entry *HashMapEntry) IsBlank() bool { return entry.r == -1 }

type HashMap struct {
	size  int
	table []HashMapEntry
}

func (hm *HashMap) loadFactor() float32 {
	return float32(hm.size) / float32(len(hm.table))
}

func (hm *HashMap) make(capacity int) {
	hm.size = 0
	hm.table = make([]HashMapEntry, capacity)
	for i := range hm.table {
		hm.table[i] = blankHashMapEntry()
	}
}

func (hm *HashMap) rehash() {
	oldTable := hm.table

	hm.make(int(1.5 * float32(len(hm.table))))

	for i := range oldTable {
		oldSlot := &oldTable[i]
		if !oldSlot.IsBlank() {
			for j := range hm.table {
				slot := &hm.table[hm.slot(oldSlot.r, j)]
				if slot.IsBlank() {
					*slot = HashMapEntry{r: oldSlot.r, freq: oldSlot.freq}
					break
				}
			}
			hm.size++
		}
	}
}

func (hm *HashMap) slot(r rune, iteration int) int {
	// linear probing for simplicity and cache locality
	return (hash(r) + iteration) % len(hm.table)
}

func NewHashMap() *HashMap {
	hm := new(HashMap)
	hm.make(HASH_MAP_INIT_SIZE)
	return hm
}

func (hm *HashMap) Increment(r rune) {
	for i := range hm.table {
		slot := &hm.table[hm.slot(r, i)]
		if slot.IsBlank() {
			freq := 1
			*slot = HashMapEntry{r: r, freq: freq}
			hm.size++
			if hm.loadFactor() > HASH_MAP_MAX_LOAD_FACTOR {
				hm.rehash()
			}
			return
		}
		if slot.r == r {
			slot.freq++
			return
		}
	}
	panic("should have returned")
}

// reports whether r is neither absent nor already 0 frequency
func (hm *HashMap) Decrement(r rune) bool {
	for i := range hm.table {
		slot := &hm.table[hm.slot(r, i)]
		if slot.IsBlank() {
			return false
		}
		if slot.r == r {
			slot.freq--
			return slot.freq >= 0
		}
	}
	panic("should have returned")
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hm := NewHashMap()
	for _, r := range s1 {
		hm.Increment(r)
	}

	for _, r := range s2 {
		if !hm.Decrement(r) {
			return false
		}
	}
	return true
}
