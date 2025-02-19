package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	mpBroken := make(map[rune]bool)

	for _, letter := range brokenLetters {
		mpBroken[letter] = true
	}

	split := strings.Split(text, " ")
	countBroken := len(split)
	for _, word := range split {
		fmt.Println("word")
		for _, letter := range word {
			fmt.Println("word1")
			_, ok := mpBroken[letter]
			if ok {
				fmt.Println("word2")
				countBroken--
				break
			}
		}
	}

	return countBroken
}

func main() {
	canBeTypedWords("leet code", "lt")
}