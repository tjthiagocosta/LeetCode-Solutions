func isAnagram(s string, t string) bool {
	// Check if the lengths are different
	if len(s) != len(t) {
		return false
	}
	// Create a frequency map for s
	freqMap := make(map[rune]int)

	for _, char := range s {
		freqMap[char]++
	}
	// Decrement frequency for each character in t
	for _, char := range t {
		freqMap[char]--
	}
	// Check if all frequencies are zero
	for _, freq := range freqMap {
		if freq != 0 {
			return false
		}
	}

	return true
}