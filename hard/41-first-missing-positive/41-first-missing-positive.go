func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Use index as a hash key and number sign as a presence detector
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// Step 2: Use index position as a key and inspect the negative marker
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// Case: if all elements are present in the array
	return n + 1
}