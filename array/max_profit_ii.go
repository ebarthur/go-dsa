package array

// MaxProfitII solves problem in O(n) time and O(1) space
func MaxProfitII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

// To solve this, I can maximize the profit by summing up all the "uphill" differences
//  in the stock prices.

// summing all the uphill differences work
