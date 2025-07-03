package main

import (
	"fmt"
)

func maxAbsoluteSum(nums []int) int {
    maxSum := nums[0]
    currentSum1 := nums[0]
	minSum := nums[0]
    currentSum2 := nums[0]


	for i := 1; i < len(nums); i++ {
		currentSum1 = max(nums[i], currentSum1+nums[i])
        maxSum = max(maxSum, currentSum1)
		currentSum2 = min(nums[i], currentSum2+nums[i])
        minSum = min(minSum, currentSum2)
	}

	return max(maxSum, -minSum)
}

func main() {
	nums1 := []int{1,-3,2,3,-4} // 5
	nums2 := []int{2,-5,1,-4,3,-2} // 8

	fmt.Println(maxAbsoluteSum(nums1))

	fmt.Println(maxAbsoluteSum(nums2))
}