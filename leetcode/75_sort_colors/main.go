package main

import (
	"fmt"
)

func sortColors(nums []int)  {
	count := [3]int{} // Индексы: 0, 1, 2

    for _, num := range nums {
        count[num]++
    }

	index := 0
	for val, ctn := range count {
		for i := 0; i < ctn; i++ {
			nums[index] = val
			index++
		}
	}
}

func main() {
	nums := []int{2,0,2,1,1,0}

	// nums := []int{2,0,1}

	sortColors(nums)

	fmt.Println(nums)


}