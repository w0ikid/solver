package main

import "fmt"

type Interval struct{
	stard, end int
}

func main(){
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, s, m int
		fmt.Scan(&n, &s, &m)
		intervals := make([]Interval, n)
		for j := 0; j < n; j++ {
			var l, r int
			fmt.Scan(&l, &r)
			intervals[j] = Interval{
				stard: l,
				end:   r,
			}
		}
	}
}