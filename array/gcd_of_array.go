package array

import "slices"

// Given an integer array nums, return the greatest common divisor
// of the smallest number and largest number in nums.

func NaiveFindGCD(nums []int) int {
	slices.Sort(nums) /// O(nlogn)

	return gcd(nums[0], nums[len(nums)-1])
}

// FindGCD solves the problem in O(n) time since the GCD calculation is
// negligible, O(1) space
func FindGCD(nums []int) int {
	// Find the minimum and maximum values in the array (O(n))
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Compute the GCD of the minimum and maximum values (O(log(min, max)))
	return gcd(min, max)
}

// O(log(min(a, b)))
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
