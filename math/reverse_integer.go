package math

// Reverse solves the problem in  O(n) time and O(1) space
func Reverse(x int) int {
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// Check for overflow/underflow
	if reversed > (1<<31)-1 || reversed < -(1<<31) {
		return 0
	}

	return reversed
}
