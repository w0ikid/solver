package main

func maxAscendingSum(nums []int) int {
    max := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			if j == i {
				sum += nums[j]
			} else if nums[j] > nums[j-1] {
				sum += nums[j]
			} else {
				break
			}
		}
		if sum > max {
			max = sum
		}
	}

	return max
}