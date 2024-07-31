package main

import "fmt"

func main(){
	var guest, owner, trash string
	fmt.Scan(&guest, &owner, &trash)

	trash_map := make(map[rune]int)

	for _, letter := range trash {
		trash_map[letter]++
	}

	for _, letter := range guest {
		trash_map[letter]--
	}

	for _, letter := range owner {
		trash_map[letter]--
	}

	for _, quantity := range trash_map{
		if quantity != 0 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}