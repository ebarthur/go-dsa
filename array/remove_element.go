package array

// RemoveElement removes all occurrences of val in nums in-place.
// Returns the number of elements in nums which are not equal to val.
func RemoveElement(nums []int, val int) int {
	k := 0 // Counter for elements not equal to val

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}
