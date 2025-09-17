package math

// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

// O(1)
func AddDigits(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}

	return num % 9
}

// The implementation of the AddDigits function leverages a mathematical
// property called the digital root to achieve the same result as
// recursively adding the digits of a number until only one digit remains.

// If num == 0, the result is 0.
// If num % 9 == 0 and num != 0, the result is 9.
// Otherwise, the result is num % 9.

// This property works because of the relationship between a number and its modulo 9 value.
