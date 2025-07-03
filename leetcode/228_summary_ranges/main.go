package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	
	k := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] + 1 {
			continue
		} else {
			if k == i - 1 || k == i {
				result = append(result, strconv.Itoa(nums[k]))
				k = i
			} else {
				str := strconv.Itoa(nums[k]) + "->" + strconv.Itoa(nums[i-1])
				result = append(result, str)
				k = i
			}
		}
	}

	return result
}

func main() {
	nums1 := []int{0,1,2,4,5,7}
	nums2 := []int{0,2,3,5,7,8,9}

	fmt.Println(summaryRanges(nums1))
	fmt.Println(summaryRanges(nums2))
}