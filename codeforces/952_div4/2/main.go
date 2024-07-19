package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b string
		fmt.Scan(&a, &b)

		ar := []rune(a)
		br := []rune(b)

		ar[0], br[0] = br[0], ar[0]

		a = string(ar)
		b = string(br)

		fmt.Println(a, b)
	}
}