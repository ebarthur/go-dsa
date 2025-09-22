package greedy

// JumpGameII solves the problem of finding the minimum number of jumps to reach the last index
func JumpGameII(nums []int) int {
	jumps := 0      // Number of jumps made so far
	currentEnd := 0 // End of the current jump range
	farthest := 0   // Farthest index reachable so far

	for i, n := range nums {
		// Update the farthest reachable index
		farthest = max(farthest, i+n)

		// If we reach the end of the current jump range
		if i == currentEnd {
			jumps++               // Increment the jump count
			currentEnd = farthest // Update the current jump range

			// If the current range includes the last index, we can stop
			if currentEnd >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}
