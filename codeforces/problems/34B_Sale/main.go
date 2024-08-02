package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	price := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&price[i])
	}

	sort.Ints(price)

	result := 0
	count := 0

	for i := 0; i < n && count < m; i++ {
		if price[i] < 0 {
			result += -price[i]
			count++
		}
	}

	fmt.Println(result)
}
