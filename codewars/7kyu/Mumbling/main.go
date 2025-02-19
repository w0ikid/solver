package main
// change lol
import (
	"fmt"
	"unicode"
	"strings"
)

func Accum(s string) string {
	var builder strings.Builder
	n := len(s)
	k := 0
	for _, v := range s {
		builder.WriteRune(unicode.ToUpper(v))
		for i := 0; i < k; i++ {
			builder.WriteRune(unicode.ToLower(v))
		}
		k++
		if (!(n == k)){
			builder.WriteRune('-')
		}
	}

	return builder.String()
}

func main(){
	sometext := "soFeText"

	fmt.Println(Accum(sometext))
}