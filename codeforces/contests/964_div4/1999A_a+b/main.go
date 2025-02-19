package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		sum := num / 10 + num % 10
		fmt.Println(sum)
	}
}