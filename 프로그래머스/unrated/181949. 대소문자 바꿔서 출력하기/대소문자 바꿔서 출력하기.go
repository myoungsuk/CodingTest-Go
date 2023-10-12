package main

import "fmt"
import "unicode"

func main() {
    var s1 string
	fmt.Scan(&s1)

	for _, r := range s1 {
		res := unicode.ToLower(r)
		if res == r {
			res = unicode.ToUpper(r)
		}
		fmt.Printf("%c", res)
	}
}