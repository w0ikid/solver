package main

func search(nums []int, target int) int {
	var low = 0
	var high = len(nums)
	var mid = 0
	for {
		if low >= high {
			break
		}
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
	}
	return -1
}