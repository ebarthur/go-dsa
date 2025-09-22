package array

// MajorityElement solves the problem in O(n) time and O(1) space
func MajorityElement(nums []int) int {
	mp := make(map[int]int)

	for _, num := range nums {
		mp[num]++

		if mp[num] > len(nums)/2 {
			return num
		}
	}

	return -1 // unreachable with the assumption that the majority element always exists
}
