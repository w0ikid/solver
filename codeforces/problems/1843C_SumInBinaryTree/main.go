package main

import "fmt"

func sumPath(n int64) int64{
	var sum int64 = 0
	for n > 0 {
		sum = sum + n
		n = n / 2
	}
	return sum
}

func main(){
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Scan(&n)
		fmt.Println(sumPath(n))
	}
}

