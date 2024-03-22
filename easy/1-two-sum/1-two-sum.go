func twoSum(nums []int, target int) []int {
	// Create a map to store the complements
	complementMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement
		complement := target - num

		// Check if the complement is in the map
		if index, ok := complementMap[complement]; ok {
			// If the complement is present, return the indices
			return []int{index, i}
		}

		// Store the current element and its index in the map
		complementMap[num] = i
	}

	// If no solution is found, return an empty slice
	return []int{}
}
