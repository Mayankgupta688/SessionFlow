/*

1. Working with Variadic Parameter
2. See the type of value in Variadic Parameter parameter
3. Introducing to the Slices..
4. Passing Values of different type in Variadic Parameter

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	CallRandomFunction(1, 12, 3, 4, 5, 7, 9)
}

func CallRandomFunction(data ...int) {
	fmt.Println(len(data))
	fmt.Println(reflect.TypeOf(data))
}
