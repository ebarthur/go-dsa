package oas

// TikTok 2024

// You’re given an array favCategories where each number represents how many favorite content categories a TikTok user has.

// The number of common interests between two users is defined as the greatest common divisor (GCD) of their values.

// Your task:
// Find the maximum GCD among all possible user pairs.

// Brute force O(n^2 log M)
func BFMaxSharedCategories(favCategories []int) int {
	n := len(favCategories)
	maxGcd := 0

	for i := range n {
		for j := i + 1; j < n; j++ {
			g := gcd(favCategories[i], favCategories[j])
			if g > maxGcd {
				maxGcd = g
			}
		}
	}

	return maxGcd
}

// Optimized O(M log M), where M = max element
func MaxSharedCategories(favCategories []int) int {
	// find maximum value
	max := 0
	for _, v := range favCategories {
		if v > max {
			max = v
		}
	}

	// frequency array
	freq := make([]int, max+1)
	for _, v := range favCategories {
		freq[v]++
	}

	// check divisors from high to low
	for d := max; d >= 1; d-- {
		count := 0
		for multiple := d; multiple <= max; multiple += d {
			count += freq[multiple]
			if count >= 2 {
				return d
			}
		}
	}

	return 1 // fallback (at least gcd=1 always exists if n>=2)
}

// Euclidean algo O(log(min(a,b)))
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
