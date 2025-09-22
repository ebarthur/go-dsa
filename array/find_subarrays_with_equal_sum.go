package array

// Given a 0-indexed integer array nums, determine whether there exist
// two subarrays of length 2 with equal sum. Note that the two subarrays
//  must begin at different indices.
// Return true if these subarrays exist, and false otherwise.

// FindSubarrays solves the problem in O(n) time and O(n) space
func FindSubarrays(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}

	mp := make(map[int][][]int)

	for i := 0; i < len(nums)-1; i++ {

		sum := nums[i] + nums[i+1]
		if _, exists := mp[sum]; exists {
			return true
		}
		mp[sum] = append(mp[sum], []int{nums[i], nums[i+1]})

	}
	return false
}
