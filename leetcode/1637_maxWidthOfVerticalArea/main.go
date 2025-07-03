package main

// import (
// 	"fmt"
// )

// func maxTotalReward(rewardValues []int) int {
// 	n := len(rewardValues)
// 	x := 0


// 	marked := make([]bool, n)

// 	for {
// 		found := false
// 		maxReward := -1
// 		maxIndex := -1
		
// 		for i := 0; i < n; i++ {
// 			if !marked[i] && rewardValues[i] > x {
// 				if rewardValues[i] > maxReward {
// 					maxReward = rewardValues[i]
// 					maxIndex = i
// 				}
// 				found = true
// 			}
// 		}
// 		if !found {
// 			break
// 		}
// 		x += rewardValues[maxIndex]
// 		marked[maxIndex] = true
// 	}

// 	return x
// }

// func main() {
// 	rewardValues1 := []int{1, 1, 3, 3}
// 	fmt.Println(maxTotalReward(rewardValues1)) // Вывод: 4

// 	rewardValues2 := []int{1, 6, 4, 3, 2}
// 	fmt.Println(maxTotalReward(rewardValues2)) // Вывод: 11
// }

// for i := 0; i < n; i++ {
// 	if !marked[i] && rewardValues[i] > x {
// 		if rewardValues[i] > maxReward {
// 			maxReward = rewardValues[i]
// 			maxIndex = i
// 		}
// 		found = true
// 	}
// }
