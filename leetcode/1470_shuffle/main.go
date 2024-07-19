package main

import "fmt"

func shuffle(nums []int, n int) []int {
	res := make([]int, 2 * n)
	k := 0
	for i := 0; i < 2 * n; i++ {
		if i % 2 == 0 {
			res[i] = nums[k]
			k++
		} else {
			res[i] = nums[n + k - 1]
		}
	}
	return res
}


func main(){
	fmt.Println(shuffle([]int{1,2,3,4,5,6}, 3))
	fmt.Println(0 % 2 == 0)
}

