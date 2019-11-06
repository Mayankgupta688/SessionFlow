/*

1. Understanding different types in Golang
2. Creating Value of certain type
3. Adding value of wrong type to the variable
4. Implicit Conversion posibilities
5. Explicit Type conversions "strconv"
6. Using "reflect" package to determine Types

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	userName := 10
	fmt.Println(reflect.TypeOf(userName))
}
