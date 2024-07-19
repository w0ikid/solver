package main

func differenceOfSums(n int, m int) int {
	div := 0
	non_div := 0
	for i := 1; i <= n; i++ {
		if i % m == 0{
			div = div + i
		} else {
			non_div = non_div + i
		}
	}

	return non_div - div
}