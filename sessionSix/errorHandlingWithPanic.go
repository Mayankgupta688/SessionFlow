/*

1. Creating Error Situation with Panic Keyword
2. Calling Panic in Functions...

*/

package main

import "fmt"

func main() {
	fmt.Println("Starting Application")
	ExecuteFunction()
}

func ExecuteFunction() {
	fmt.Println("Function Called...")
	panic("Program Ends...")
}
