package main

import "fmt"

func main(){
	var pi, r float64
	pi = 3.141592653

	fmt.Scan(&r)

	fmt.Println(pi*r*r)
}