package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b string
		fmt.Scan(&a, &b)

		ar := []rune(a)
		br := []rune(b)

		ar[0], br[0] = br[0], ar[0]

		a = string(ar)
		b = string(br)

		fmt.Println(a, b)
	}
}


/*
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Введите строку:")
    
    // Читаем строку до символа новой строки
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка чтения строки:", err)
        return
    }
    
    fmt.Println("Вы ввели:", input)
}
*/