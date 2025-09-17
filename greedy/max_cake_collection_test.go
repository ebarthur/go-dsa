package greedy

import "testing"

/*
TestMaxCakesCollection tests solution(s) with the following signature and problem description:

	func MaxCakesCollection(n, m int, rates []int) int

Maple wants to bake some cakes for Chocola and Vanilla.

One day, she discovers n magical cake ovens. The i-th oven bakes rates[i] cakes every second.
The cakes remain in their respective ovens until they are collected.

At the end of each second, she may teleport to any oven (including the one she is currently at)
and collect all the cakes that have accumulated in that oven up to that point.

Your task is to determine the maximum number of cakes Maple can collect in m seconds.

For example, given n=3, m=4, and rates=[1,2,3]:
- At second 1: ovens have [1,2,3] cakes. Collect from oven 3 → get 3 cakes
- At second 2: ovens have [2,4,3] cakes. Collect from oven 1 → get 2 cakes
- At second 3: ovens have [1,6,6] cakes. Collect from oven 2 → get 6 cakes
- At second 4: ovens have [2,2,9] cakes. Collect from oven 3 → get 9 cakes
Total: 3+2+6+9 = 20 cakes
*/
func TestMaxCakesCollection(t *testing.T) {
	tests := []struct {
		n        int
		m        int
		rates    []int
		maxCakes int
	}{
		{3, 4, []int{1, 2, 3}, 20},
		{3, 2, []int{1, 2, 3}, 8},
		{1, 1000, []int{100000}, 100000000},
	}

	for i, test := range tests {
		got := MaxCakesCollection(test.n, test.m, test.rates)
		if got != test.maxCakes {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.maxCakes, got)
		}
	}
}
