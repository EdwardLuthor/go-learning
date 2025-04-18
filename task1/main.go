package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1+num2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Printf("Input: nums = %v, target = %d, Output: %v\n", nums1, target1, result1)

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("Input: nums = %v, target = %d, Output: %v\n", nums2, target2, result2)

	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Printf("Input: nums = %v, target = %d, Output: %v\n", nums3, target3, result3)

	nums4 := []int{3, 3}
	target4 := 9
	result4 := twoSum(nums4, target4)
	fmt.Printf("Input: nums = %v, target = %d, Output: %v\n", nums4, target4, result4)
}
