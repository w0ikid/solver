package main

import "fmt"

func subsetXORSum(nums []int) int {
    return subsets(nums, 0, 0)
}

func subsets(nums []int, i, curr int) int {
    if i >= len(nums) {
        return curr
    }
    return subsets(nums, i+1, nums[i]^curr) + subsets(nums, i+1, curr)
}

func main() {
	nums := []int{5, 1, 6}
	// fmt.Println(subsets(nums))

	fmt.Println(subsetXORSum(nums))
}