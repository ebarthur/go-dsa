package array

// MergeSortedArrays merges nums2 into nums1 in sorted order.
// nums1 has enough space to hold all elements from nums1 and nums2.
func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	// Start merging from the end of nums1 and nums2
	for p, p1, p2 := m+n-1, m-1, n-1; p2 >= 0; p-- {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}
}
