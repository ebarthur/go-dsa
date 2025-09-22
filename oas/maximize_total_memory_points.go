package oas

import (
	"slices"
)

// Greedy sorting problem
// Amazon OA

// You are given n chapters. Each chapter has memory[i] points (positive or negative).
// If you read chapters in some order, the points you get are the prefix sums of that order.
// Your task is to choose the order of chapters that maximizes the total points.

// Input:
// memory = [3, 4, 5]
//
// Output:
// 26

// MaximizeTotalMemoryPoint solves the problem in O(nlogn)
func MaximizeTotalMemoryPoint(points []int) int {
	slices.Sort(points)
	slices.Reverse(points)

	total, prefix := 0, 0
	for _, num := range points {
		prefix += num
		total += prefix
	}

	return total
}

// The optimal order is simply sorting the array in descending order.
// So to maximize total memory points, Place larger numbers earlier,
// since they get multiplied more times.

// for [3,4,5], sorting it becomes [5,4,3],
// first iteration: [5,4,3]
// second iteration: [5,9,3]
// third iteration: [5,9,12]

// Answer: 5 + 9 + 12 = 26
