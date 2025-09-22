package strings

// Bruteforce
// LongestCommonPrefix solves the problem in O(n^2) time and O(1) space
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	// Input: strs = ["flower","flow","flight"]

	// Compare the prefix with each string in the array
	for _, str := range strs[1:] {
		// If the current input is an empty string, return ""
		if len(str) == 0 {
			return ""
		}

		// Reduce the prefix until it matches the current string
		for len(prefix) > 0 && len(str) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		// If the prefix becomes empty, return ""
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// Helper function to check if a string starts with a given prefix
func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}

// [] TODO: Optimized
