package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for k >= 0 {
		if j < 0 {
			break
		}

		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	// test1
	// nums1 := []int{1,2,3,0,0,0}
	// nums2 := []int{2,5,6}

	// m := 3
	// n := 3

	// test2
	// nums1 := []int{1}
	// nums2 := []int{}

	// m := 1
	// n := 0

	// test 3
	nums1 := []int{0}
	nums2 := []int{1}

	m := 0
	n := 1


	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}