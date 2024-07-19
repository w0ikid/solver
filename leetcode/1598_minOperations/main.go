package main

func minOperations(logs []string) int {
	min := 0

	for _, dir := range logs {
		if dir == "../" && min > 0 {
			min--
		} else if dir == "./" {
			continue
		} else {
			min++
		}
	}

	return min
}
