package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        var s string
        fmt.Scan(&s)
        
        x, y := 0, 0
        found := false
        for _, v := range s {
            switch v {
            case 'U':
                y++
            case 'D':
                y--
            case 'L':
                x--
            case 'R':
                x++
            }
            if x == 1 && y == 1 {
                found = true
                break
            }
        }
        if found {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
