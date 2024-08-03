package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	fmt.Scan(&n)
	segments := make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&segments[i])
	}

	sort.Ints(segments)

	for i := 1; i < n-1; i++ {
		if (segments[i-1] + segments[i] > segments[i+1]){
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}