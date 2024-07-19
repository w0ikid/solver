package main

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scoreOfString(s string) int {
	result := 0
	for i := 1; i < len(s); i++ {
		result = result + absInt(int(s[i - 1]) - int(s[i]))
	}
	return result
}