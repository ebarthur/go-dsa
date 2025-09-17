package array

import "fmt"

func ArrayConcepts() {

	// declaration and initialization
	arr := []int32{1, 2, 3, 4, 5}

	fmt.Printf("Array elements: %v\n", arr)

	// Accessing elements --- O(1)
	fmt.Printf("Element at index 2: %d\n", arr[2])

	// Modifying elements --- O(1)
	arr[2] = 6
	fmt.Printf("After modifying element at pos 2: %v\n", arr)

	// Delete elements at position 2 --- O(n)
	index := 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Printf("After deleting element at pos 2: %v", arr)

}
