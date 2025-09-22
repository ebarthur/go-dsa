package array

// SortArray sorts an array of integers in ascending order using Merge Sort.
// O(nlogn) time
func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums // Base case: a single element is already sorted.
	}

	// Divide the array into two halves.
	mid := len(nums) / 2
	left := SortArray(nums[:mid])
	right := SortArray(nums[mid:])

	// Merge the two sorted halves.
	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from both slices and append the smaller one.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice.
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements from the right slice.
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
