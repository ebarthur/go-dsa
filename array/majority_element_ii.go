package array

// MajorityElementII solves the porblem in O(n) time and O(n) space
func MajorityElementII(nums []int) []int {
	var output []int
	mp := make(map[int]int)
	seen := make(map[int]bool)

	for _, num := range nums {
		mp[num]++

		if mp[num] > len(nums)/3 && !seen[num] {
			output = append(output, num)
			seen[num] = true
		}
	}

	return output
}
