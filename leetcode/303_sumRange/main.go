package main

import "fmt"

type NumArray struct {
    array []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
    	array: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := 0; i < right - left; i++ {
		sum = sum + this.array[left + i] 
	}
	return sum
}

func main(){
	someNums := [5]int{1, 2, 3, 4, 5}

	NumArray_create := Constructor(someNums[:])
	NumArray_create.SumRange(2, 5)
	fmt.Println(NumArray_create.SumRange(2, 5))
}