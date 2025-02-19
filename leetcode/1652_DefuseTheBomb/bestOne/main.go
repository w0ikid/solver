package main

func decrypt(code []int, k int) []int {
    n := len(code)
    res := make([]int, n)
    if k == 0 {
        return res
    }
    for i := range n {
        if k > 0 {
            for j:=i+1; j<i+k+1; j++ {
                res[i]+=code[j%n]
				/*
				[1, 2, 3, 4, 5]
				j = 

				*/
            } 
        } else {
            for j:=i+k; j<i; j++ {
                res[i]+=code[(j+n)%n]
            }
        }
    }
    return res
}