/*

1. Passing Values to Function by Reference
2. Passing String to Functions by Reference

*/

package main

import "fmt"

func main() {
	inputValue := 10
	sampleString := "Hello..."
	OtherFunction(&inputValue, &sampleString)
	fmt.Println(inputValue)
	fmt.Println(sampleString)
}

func OtherFunction(inputValue *int, sampleString *string) {
	*sampleString = "Some Other Data"
	*inputValue = 20
}
