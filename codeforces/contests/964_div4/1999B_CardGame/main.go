package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	results := make([]int, t)

	for i := 0; i < t; i++ {
		var a1, a2, b1, b2 int
		fmt.Scan(&a1, &a2, &b1, &b2)

		rounds := [][4]int{
			{a1, b1, a2, b2},
			{a1, b2, a2, b1},
			{a2, b1, a1, b2},
			{a2, b2, a1, b1},
		}

		winCount := 0

		for _, round := range rounds {
			sunilWin := 0
			slavekWin := 0

			if round[0] > round[1] {
				sunilWin++
			} else if round[0] < round[1] {
				slavekWin++
			}

			if round[2] > round[3] {
				sunilWin++
			} else if round[2] < round[3] {
				slavekWin++
			}

			if sunilWin > slavekWin {
				winCount++
			}
		}

		results[i] = winCount
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
