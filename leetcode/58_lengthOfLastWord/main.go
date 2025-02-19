package main

import (
	"fmt"
	"strings"
)

// func lengthOfLastWord(s string) int {
// 	data := strings.Fields(s)
// 	return len(data[len(data) - 1])
// }

// without strings.Fields()
func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main(){
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  daniaasdfasdf"))
}
