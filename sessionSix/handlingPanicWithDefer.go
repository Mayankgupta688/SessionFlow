/*

1. Defer Function called after Panic

*/

package main

import "fmt"

func main() {
	ExecuteFunction()
}

func ExecuteFunction() {
	defer HandleError()
	defer HandleErrorOther()

	panic("Function Executed Error Condition...")
}

func HandleError() {
	fmt.Println("Defer Function Executed after Panic...")
}

func HandleErrorOther() {
	fmt.Println("Defer Function Executed after Panic Other...")
}
