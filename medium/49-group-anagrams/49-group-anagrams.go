package main

func groupAnagramsYour(strs []string) [][]string {
	// Create a hashmap to store the anagrams
	// The key is a byte array of length 26 representing the count of each character
	// The value is a slice of strings that are anagrams
	hashMap := make(map[[26]byte][]string)

	// Iterate over each string in the input slice
	for _, str := range strs {
		// Create a byte array of length 26 to store the count of each character
		var count [26]byte

		// Iterate over each character in the string
		for _, char := range str {
			// Increment the count of the character in the count array
			// The index is calculated by subtracting 'a' from the character
			// This assumes the input strings contain only lowercase letters
			count[char-'a']++
		}

		// Add the current string to the slice of anagrams in the hashmap
		// The key is the count array, which is the same for all anagrams
		hashMap[count] = append(hashMap[count], str)
	}

	// Create a result slice to store the groups of anagrams
	result := make([][]string, 0, len(hashMap))

	// Iterate over the values (slices of anagrams) in the hashmap
	for _, value := range hashMap {
		// Append each slice of anagrams to the result slice
		result = append(result, value)
	}

	// Return the result slice containing the groups of anagrams
	return result
}
