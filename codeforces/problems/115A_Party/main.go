package main

import (
	"fmt"
)


// Функция для поиска глубины дерева
func findDepth(employee int, managers []int, memo map[int]int) int {
	// Проверяем, есть ли уже вычисленная глубина для этого сотрудника
	if depth, found := memo[employee]; found {
		return depth
	}

	// Находим непосредственного начальника этого сотрудника
	manager := managers[employee]
	// Если начальника нет, глубина равна 1
	if manager == -1 {
		memo[employee] = 1
		return 1
	}

	// Иначе глубина равна 1 + глубина начальника
	depth := 1 + findDepth(manager-1, managers, memo)
	memo[employee] = depth
	return depth
}

func main() {
	var n int
	fmt.Scan(&n)

	managers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&managers[i])
	}

	memo := make(map[int]int)
	maxDepth := 0
	for i := 0; i < n; i++ {
		depth := findDepth(i, managers, memo)
		if depth > maxDepth {
			maxDepth = depth
		}
	}

	fmt.Println(maxDepth)
}
