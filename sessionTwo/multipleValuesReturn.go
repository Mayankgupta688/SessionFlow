/*

1. Create Function having multiple return value
2. Error when the other value is not received...
3. Avoiding specific value returned from the function
4. Using "_"to avoid values...

*/

package main

import "fmt"

func main() {
	firstVal, _ := functionWithMultipleParam()
	fmt.Println(firstVal)
}

func functionWithMultipleParam() (int, int) {
	return 1, 2
}
