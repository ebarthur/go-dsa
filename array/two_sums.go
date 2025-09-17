package array

// bruteforce solution O(n^2)
// func twoSum(nums []int, target int) []int {
//     // bruteforce O(n^2)
//     for i := 0; i <= len(nums) -1; i++ {
//         for j := i+1; j <= len(nums) -1; j++ {
//             if nums[i] + nums[j] == target {
//                 return []int{i,j}
//             }
//         }
//     }
//     return []int{-1,-1}
// }

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	// loop once O(n)
	for i, num := range nums {
		diff := target - num

		// map lookup and insertion O(1)
		if j, exists := mp[diff]; exists {
			return []int{j, i}
		}

		mp[num] = i

	}
	return []int{-1, -1}
}

// Overall time complexity - O(n)
// Overall space complexity - O(n) - mp can store up to n key-value pairs in worst case