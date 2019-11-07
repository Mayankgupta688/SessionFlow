/*

1. Creating Custom Functions returning Errors using "errors.New"
2. Error Last Approach in Golang
3. Value passed to "error.New" is received as error value...

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	returnData, err := ThrowErrorFunc()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(returnData)
	}
}

func ThrowErrorFunc() (int, error) {
	return 0, errors.New("Invalid Operation Data...")
}
