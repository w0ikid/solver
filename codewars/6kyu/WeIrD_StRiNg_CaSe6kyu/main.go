package main

import (
	"fmt"
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	var builder strings.Builder
	k := 0
	for _, char := range str {
		if k % 2 == 0 {
			builder.WriteRune(unicode.ToUpper(char))
		} else {
			builder.WriteRune(unicode.ToLower(char))
		}
		k++
	}
	return builder.String()
}

func main(){
	sometext := "Lorem ipsum"
	someChar := ' '
	fmt.Println(unicode.ToUpper(someChar))
	for _, v := range sometext {
		fmt.Println(v)
	}
	fmt.Println(toWeirdCase(sometext))
}