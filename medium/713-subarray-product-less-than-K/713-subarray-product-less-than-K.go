package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	n := len(nums)
	product := 1
	left, right := 0, 0 // Left and right pointers for the sliding window
	count := 0

	// Sliding window approach
	for right < n {
		product *= nums[right] // Update product by multiplying with the current element

		// Shrink the window from the left side until the product is less than k
		for product >= k {
			product /= nums[left]
			left++
		}

		// Count the number of valid subarrays ending at the current right pointer
		count += right - left + 1

		right++
	}

	return count
}
