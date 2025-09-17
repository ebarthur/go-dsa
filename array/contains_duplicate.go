package array

// Given an integer array nums, return true if any value appears at
// least twice in the array, and return false if every element is
// distinct.

// ContainsDuplicate solves the problem in O(n) time and O(n) space
func ContainsDuplicate(nums []int) bool {
	mp := make(map[int]bool)

	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = true
	}

	return false
}
