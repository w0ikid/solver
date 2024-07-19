package main

import (
	"strings"
)

func numJewelsInStones(jewels string, stones string) int {
    count := 0
	for _, item := range jewels{
		count = count + strings.Count(stones, string(item))
	}
	return count
}