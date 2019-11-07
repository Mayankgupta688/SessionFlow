/*

1. Working this "_" identifier for range function
2. Creating Custom Function returning Miultiple values
3. Using "_" for import libraries

*/

package main

import (
	"fmt"
	_ "os"
)

func main() {
	integerArray := [5]int{1, 2, 3, 4, 10}

	for index, val := range integerArray {
		fmt.Println(index, " ", val)
	}

	_, otherData := GetData()

	fmt.Println(otherData)
}

func GetData() (int, int) {
	return 1, 4
}
