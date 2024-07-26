package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	coins := make([]int, n)
	totalAmount := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&coins[i])
		totalAmount += coins[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	var takenAmount, minCoins int

	for i := 0; i < n; i++ {
		takenAmount += coins[i]
		totalAmount -= coins[i]
		minCoins++
		if takenAmount > totalAmount {
			break
		}
	}

	fmt.Println(minCoins)
}
