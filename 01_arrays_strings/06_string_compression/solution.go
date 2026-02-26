package stringcompression

import "strconv"

func StringCompression(s string) string {
	if len(s) < 2 {
		return s
	}

	compressed := make([]rune, 0, len(s))

	curr := rune(-1)
	runLen := 0
	writeCurr := func() {
		if curr != -1 {
			compressed = append(compressed, curr)
			compressed = append(compressed, []rune(strconv.Itoa(runLen))...)
		}
	}
	
	for _, r := range s {
		if r == curr {
			runLen++
			continue
		}

		writeCurr()
		if len(compressed) >= len(s) {
			return s
		}

		curr = r
		runLen = 1
	}

	writeCurr()
	if len(compressed) >= len(s) {
		return s
	}

	return string(compressed)
}
