package main

import "fmt"

func toBinary(n int) string {
    result := ""
	for n != 0 {
		if n % 2 == 0{
			result = "0" + result
		} else {
			result = "1" + result
		}
		n = n / 2
	}
	return result
}

func minBitFlips(start int, goal int) int {
    count := 0
	first := toBinary(start)
	second := toBinary(goal)
	if len(second) < len(first){
		for i := 0; i < len(second); i++ {
			second = "0" + second
		}
	} else {
		for i := 0; i < len(first); i++ {
			first = "0" + first
		}
	}
	for i := 0; i < len(second); i++ {
		if first[i] == second[i]{
			count++
		}
	}
	return count
}

func main(){
	fmt.Println(toBinary(7))
	fmt.Println(toBinary(10))
	fmt.Println(minBitFlips(10, 7))
}