package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    var result [2]int

	n := len(nums)

	// 00 01 02 03 04
	// 10 11 12 13 14
	// 20 21 22 23 24
	// 30 31 32 33 34
	// 40 41 42 43 44


	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if (nums[i] + nums[j]) == target {
				result[0] = i
				result[1] = j
				return result[:]
			}
		}
	}

	return result[:]
}

func main() {
	nums := []int{2,7,11,15}
	target := 9

	nums1 := []int{3,2,4}
	target1 := 6

	nums2 := []int{3,3}
	target2 := 6

	nums3 := []int{2,5,5,11}
	target3 := 10

	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum(nums1, target1))
	fmt.Println(twoSum(nums2, target2))

	fmt.Println(twoSum(nums3, target3))


	shit()
}

// 01 02 03 04
// 12 13 14
// 23 24
// 34

func shit() {
	for i := 0; i < 5; i++ {
		for j := i+1; j < 5; j++ {
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Printf("\n")
	}
}