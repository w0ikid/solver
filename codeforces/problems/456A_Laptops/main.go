package main

import (
	"fmt"
	"sort"
)

type Laptop struct {
	Price int
	Power int
}

func main(){
	var n int
	fmt.Scan(&n)

	laptops := make([]Laptop, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&laptops[i].Price, &laptops[i].Power)
	}
	
	sort.Slice(laptops, func(i, j int) bool {
		return laptops[i].Power < laptops[j].Power
	})

	for i := 1; i < n; i++ {
		if laptops[i].Price < laptops[i - 1].Price {
			fmt.Println("Happy Alex")
			return
		}
	}
	fmt.Println("Poor Alex")
}