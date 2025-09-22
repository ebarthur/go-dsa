package array

// Given an integer array nums, find a subarray that has the
// largest product, and return the product.

// MaxProduct solves problem in O(n) time and O(1) space
func MaxProduct(nums []int) int {
	maxProduct := nums[0]
	currentMin := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			// Swap currentMin and currentMax when we encounter negative
			currentMin, currentMax = currentMax, currentMin
		}

		currentMin = min(nums[i], currentMin*nums[i])
		currentMax = max(nums[i], currentMax*nums[i])

		maxProduct = max(maxProduct, currentMax)

	}
	return maxProduct
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
