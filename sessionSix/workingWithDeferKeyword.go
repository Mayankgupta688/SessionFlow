/*

1. Creating Defered Functions
2. Adding Multiple Defer to the Function (Lifo)

*/

package main

import "fmt"

func main() {

	defer DeferFunction()
	defer OtherFunction()
	fmt.Println("This is the normal Execution flow")
}

func DeferFunction() {
	fmt.Println("This is Deferred Function...")
}

func OtherFunction() {
	fmt.Println("This is Deferred Other Function...")
}
