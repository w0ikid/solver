package main

import "fmt"

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

	

}