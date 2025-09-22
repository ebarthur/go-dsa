package greedy

// JumpGame solves the problem in O(n) time and O(1) space.
func JumpGame(jumps []int) bool {
	if len(jumps) == 0 {
		return true // Edge case: empty array
	}

	maxReachableDistance := 0 // Initialize the farthest reachable index

	for i := range jumps {
		// If the current index is beyond the maximum reachable distance, return false
		if i > maxReachableDistance {
			return false
		}

		// Update the maximum reachable distance
		maxReachableDistance = max(maxReachableDistance, i+jumps[i])
	}

	// Check if the last index is reachable
	return maxReachableDistance >= len(jumps)-1
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

