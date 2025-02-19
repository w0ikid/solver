package main

import "fmt"

func main(){
	var n int;
	fmt.Scan(&n)

	var years int;
	var month int;
	var days int;

	years = n / 365
	month = (n % 365) / 30
	days = (n % 365) % 30
}