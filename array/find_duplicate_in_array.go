package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
func FindDuplicate(list []int) int {
	for _, num := range list {
		itemIndex := abs(num) - 1

		if list[itemIndex] < 0 {
			return abs(num)
		}

		list[itemIndex] *= -1
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
