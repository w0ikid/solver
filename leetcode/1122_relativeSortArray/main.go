package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
    // Step 1: Create a counting array for the range of numbers (0 to 1000)
    count := make([]int, 1001)

    // Step 2: Populate the counting array with the frequency of each number in arr1
    for _, num := range arr1 {
        count[num]++
    }

    // Step 3: Create a result array
    res := make([]int, 0, len(arr1))

    // Step 4: Add numbers from arr2 to the result array according to their frequency
    for _, num := range arr2 {
        for count[num] > 0 {
            res = append(res, num)
            count[num]--
        }
    }

    // Step 5: Add remaining numbers to the result array in ascending order
    for num, freq := range count {
        for freq > 0 {
            res = append(res, num)
            freq--
        }
    }

    return res
}
