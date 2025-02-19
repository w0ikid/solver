package main

import "fmt"

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	// (i + 1 + j) % len(code)
	lenCode := len(code)
	if k == 0 {
		return result
	} else if k > 0 {
		for i := 0; i < lenCode; i++ {
			sum := 0
			for j := 0; j < k; j++ {
				sum  = sum + code[(i+1+j)%lenCode]
			}
			result[i] = sum
			sum = 0
		}
	} else if k < 0 {
		// (len+k+i+j)%len
		for i := 0; i < lenCode; i++ {
			sum := 0
			for j := 0; j < -k; j++ {
				sum  = sum + code[(lenCode+k+i+j)%lenCode]
			}
			result[i] = sum
			sum = 0
		}
	}

	return result
}

func main(){
	fmt.Println("Test :>")

	array := []int{1, 2, 3, 4, 5}
	array2 := []int{1, 2, 3, 4, 5}
	k := 3
	test := decrypt(array, k)
	test2 := decrypt(array2, -3)
	test3 := decrypt(array, 0)
	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(test3)
}

/*
[1, 2, 3, 4, 5]
k = -3
len - k = 2
(len + k + i + j) % len
(5 - 3 + 0 + 0) % 5 = 2
(5 - 3 + 0 + 1) % 3 = 3
(5 - 3 + 0 + 2) % 3 = 4
---
(5 - 3 + 1 + 0) % 5 = 3
(5 - 3 + 1 + 1) % 5 = 4
(5 - 3)

k = -3
[12, 10, 8, 6, 9]
k = 3
[9, 12, 10, 8, 6]
-----------------
k = -2
]
*/