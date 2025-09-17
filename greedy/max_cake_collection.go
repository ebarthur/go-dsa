package greedy

// greedy solution (Not optimal)

func GreedyMaxCakesCollection(n, m int, rates []int) int {
	if m == 0 || n == 0 {
		return 0
	}

	acc := make([]int, n)
	total := 0

	for sec := 1; sec <= m; sec++ {
		for i := range n {
			acc[i] += rates[i]
		}

		max := 0
		for i := range n {
			if acc[i] > acc[max] {
				max = i
			}
		}

		total += acc[max]
		acc[max] = 0
	}

	return total
}

// dp
func MaxCakesCollection(n, m int, rates []int) int {
	if m == 0 || n == 0 {
		return 0
	}

	memo := make(map[string]int)

	initial := make([]int, n)

	return dp(0, m, rates, initial, memo)
}

func dp(current, total int, rates []int, ovens []int, memo map[string]int) int {
	if current == total {
		return 0
	}

	key := key(current, ovens)
	if val, exists := memo[key]; exists {
		return val
	}

	n := len(rates)

	newOvens := make([]int, n)
	for i := range n {
		newOvens[i] = ovens[i] + rates[i]
	}

	max := 0

	for i := range n {
		currentOvenCakes := newOvens[i]

		nextOvens := make([]int, n)
		copy(nextOvens, newOvens)
		nextOvens[i] = 0

		remaining := dp(current+1, total, rates, nextOvens, memo)

		totalCakes := currentOvenCakes + remaining
		if totalCakes > max {
			max = totalCakes
		}
	}

	memo[key] = max
	return max
}

func key(sec int, ovens []int) string {
	key := string(rune(sec)) + ":"
	for _, count := range ovens {
		key += string(rune(count)) + ","
	}
	return key
}
