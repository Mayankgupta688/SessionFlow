/*

1. Creating Package Level Variables
2. Function calling and Updating the Package Level Variables
3. Creating Function level variable with shorthand notation
4. Overriding the Package Level Variables
5. Defining multiple "var" with var () and value assignment
6. Using "os.Getenv" in the value assignment
7. Using "init()" function
8. Function Level Variable must be used, no restriction on Package Level for error
9. Working with constant keyword "const"

*/

package main

import (
	"fmt"
)

var Name string

var pathData string

func init() {
	pathData = "Sample"
}

func main() {
	Name := "Mayank"
	UpdateName()
	fmt.Println(Name)
	fmt.Println(pathData)
}

func UpdateName() {
	Name = "Änshul"
	fmt.Println(Name)
}
