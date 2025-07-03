package main

import "fmt"

func removeDuplicates(nums []int) int {
	index := 1
	
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func main() {
	nums1 := []int{1,1,2}

	nums2 := []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
}