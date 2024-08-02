package main

func findTheDifference(s string, t string) byte {
	difference := make(map[rune]int)
	for _, letter := range s {
		difference[letter]++
	}
	for _, letter := range t {
		difference[letter]--
	}
	for value, count := range difference{
		if count < 0 {
			return byte(value)
		}
	}
	return 0
}