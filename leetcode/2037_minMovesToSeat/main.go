package main

import (
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	count := 0

	for i := 0; i < len(seats); i++ {
		count += int(math.Abs(float64(seats[i] - students[i])))
	}
	return count
}