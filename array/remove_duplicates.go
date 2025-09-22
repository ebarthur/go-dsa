package array

// RemoveDuplicates solves the problem in O(n) time and O(1) space
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Pointer to place the next valid element
	j := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		// Allow at most two of each element.
		//
		// For problems where we need to remove duplicates
		// such that each distinct item exists only once, 
		// the conditional below will be `count <= 1`
		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
