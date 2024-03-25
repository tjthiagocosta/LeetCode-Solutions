numSet := make(map[int]bool)
for _, num := range nums {
	if numSet[num] {
		return true
	}
	numSet[num] = true
}
return false