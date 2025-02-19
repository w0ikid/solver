package main

import (
	"strings"
)

func Accum(s string) string {
	arr := make([]string, len(s))

	

	for i := 0; i < len(s); i++ {
		arr[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}

	return strings.Join(arr, "-")
}