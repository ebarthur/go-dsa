package array

// BFProductExceptSelf solves the problem in O(n^2) time and O(n) space
func BFProductExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums)) // O(n)

	for i := 0; i <= len(nums)-1; i++ {
		product := 1

		for j := 0; j <= len(nums)-1; j++ {
			if i != j {
				product *= nums[j]
			}
		}
		result = append(result, product)

	}
	return result
}

// Let's try again without the nested loop
// we would iterate from start to :n and then do from n+1

// ProductExceptSelf solves the problem in O(n) time and O(n) space
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := make([]int, n)
	right := make([]int, n)

	left[0] = 1
	right[n-1] = 1

	// O(n)
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// O(m)
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// O(k)
	for i := range n {
		result[i] = left[i] * right[i]
	}
	return result
}

// OptimalProductExceptSelf solves the problem in O(n) time and O(n) space
// It creates less arrays as it calculates the prefix and suffix products
// directly

// This is very similar to the one just above without the need for extra (left
// and right) arrays
func OptimalProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate prefix products directly in the result array
	prefix := 1
	for i := range n {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Multiply by suffix products directly in the result array
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
