package main

import (
	"fmt"
	"sort"
)

type Dragon struct {
	Power int
	Bonus int
}

func main() {
	var s, n int
	fmt.Scan(&s, &n)

	dragons := make([]Dragon, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&dragons[i].Power, &dragons[i].Bonus)
	}
	
	sort.Slice(dragons, func(i, j int) bool {
		return dragons[i].Power < dragons[j].Power
	})

	for _, dragon := range dragons {
		if s > dragon.Power {
			s += dragon.Bonus
		} else {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
