package array

// RotateKSteps solves the problem in O(n) time and O(1) space.
func RotateKSteps(list []int, k int) {
	n := len(list)
	k = k % n

	// Reverse the entire array
	ReverseInPlace(list, 0, n-1)

	// Reverse the first k elements
	ReverseInPlace(list, 0, k-1)

	// Reverse the remaining elements
	ReverseInPlace(list, k, n-1)
}
