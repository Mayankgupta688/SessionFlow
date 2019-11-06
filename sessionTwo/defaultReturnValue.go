/*

1. Function returning parameter without "return"
2. Function returning without return keyword.. Gives error

*/

package main

import "fmt"

func main() {
	a, b := callFunctionWithoutReturn()
	fmt.Println(a)
	fmt.Println(b)
}

func callFunctionWithoutReturn() (first int, second int) {
	first = 1
	second = 2
	return
}
