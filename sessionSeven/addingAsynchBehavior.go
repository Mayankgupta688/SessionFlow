/*

1. Creating Simple Synchronous Function

*/

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting Main Thread...")

	go IterateData()

	go IterateOtherData()

	time.Sleep(5 * time.Second)

	fmt.Println("Stopping Main Function...")
}

func IterateData() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func IterateOtherData() {
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
}
