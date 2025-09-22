package array

import (
	"sort"
)

// ThreeSum solves the problem in  O(n^2) time and O(n) space
// same as zero_sum_triplets.go
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums) // Step 1: Sort the array
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				// Found a triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for the second and third numbers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers
				left++
				right--
			} else if sum < target {
				left++ // Increase the sum
			} else {
				right-- // Decrease the sum
			}
		}
	}

	return result
}

