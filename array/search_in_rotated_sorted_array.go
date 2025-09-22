package array

// Search finds the target in a rotated sorted array in O(log n) time
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		// Check if mid is the target
		if nums[mid] == target {
			return mid
		}

		// Determine which half is sorted
		if nums[left] <= nums[mid] { // Left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Target is in the left half
			} else {
				left = mid + 1 // Target is in the right half
			}
		} else { // Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Target is in the right half
			} else {
				right = mid - 1 // Target is in the left half
			}
		}
	}

	return -1 // Target not found
}
