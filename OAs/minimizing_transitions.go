package oas

// TikTok 2024

// Problem: Minimize Transitions in Binary Array

// You are given a binary array bits of length n where:

// 0 represents standard content,

// 1 represents priority content.

// A transition is defined as an index i (0 <= i < n-1) where bits[i] != bits[i+1].

// Your task is to minimize the number of transitions in the array by performing swap operations. In one operation, you may swap the values at any two distinct indices.

// Return the minimum number of swap operations required so that the number of transitions in the array is minimized.

// [1, 0, 1, 0, 1]
// [0, 0, 1, 1, 1]
// A binary array can always be arranged optimally
//
//	by just grouping all 0s first, then 1s.
//
// That’s basically the Dutch National Flag problem,
// which can be solved in O(n) time with just a single pass.
func GetOptimalContentStorage(bits []int) int {
	// Count how many zeros there should be
	zeros := 0
	for _, b := range bits {
		if b == 0 {
			zeros++
		}
	}

	// Count how many 1s are wrongly placed in the zero zone
	misplaced := 0
	for i := 0; i < zeros; i++ {
		if bits[i] == 1 {
			misplaced++
		}
	}

	return misplaced
}
