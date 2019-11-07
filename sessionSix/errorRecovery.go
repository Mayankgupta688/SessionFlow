/*

1. "recovery" functions in defer
2. Function returned to calling function after "panic" occurs and recovered..

*/

package main

import "fmt"

func main() {
	ExecuteError()
	fmt.Println("Error Recovered...")
}

func ExecuteError() {
	defer HandleError()
	panic("Error Occured...")
}

func HandleError() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
}
