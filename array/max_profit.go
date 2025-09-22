package array

import "math"

// MaxProfit solves the problem in time O(n) and space O(1)
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt // Initialize to the maximum possible integer
	maxProfit := 0

	for _, price := range prices {
		minPrice = min(minPrice, price)

		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

// [7 1 5 3 6 4]
// Bruteforce O(n^2)
func BFMaxProfit(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices)-1; i++ {
		current := prices[i]
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - current
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
