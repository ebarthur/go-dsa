package array

// Given the sorted rotated array nums of unique elements, return the 
// minimum element of this array.

// NOTE: You must write an algorithm that runs in O(log n) time.

// FindMin solves the problem in O(logn) time 
func FindMin(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := (left + right) / 2

        // If mid element > right element,
        // the min is in the right half
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            // Otherwise, min is in left half (including mid)
            right = mid
        }
    }

    // left == right → minimum element
    return nums[left]
}

// Using Binary search ensures that meet contraints