/*

1. Creating WaitGroup with "sync.WaitGroup"
2. Adding Wait using .Wait() functions
3. Adding WaitGroup to main using ".Add" function
4. Marking WaitGroup with ".Done"

*/

package main

import (
	"fmt"
	"sync"
)

var syncLock sync.WaitGroup

func main() {

	fmt.Println("Main Thread Started...")

	syncLock.Add(2)

	go CallAsyncFunction()
	go CallAsyncFunction()

	syncLock.Wait()

	fmt.Println("Main Thread Executed...")

}

func CallAsyncFunction() {
	fmt.Println("This is Seperate Goroutine...")
	go EmbeddedGoRoutine()
}

func EmbeddedGoRoutine() {
	fmt.Println("Embedded Goroutine Executed...")
	syncLock.Done()
}
