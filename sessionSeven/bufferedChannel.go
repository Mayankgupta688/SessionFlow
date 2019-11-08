/*

1. Create a simple Channel in Golang...
2. Add Data to the Channel..
3. See the deadlocks
4. Recieve the channel data..
5. See the deadlock again
6. Add Buffered Channels....

*/

package main

import "fmt"

func main() {
	fmt.Println("Main Execution Started..")
	channelImplementation := make(chan int, 1)
	channelImplementation <- 10
	<-channelImplementation
	fmt.Println("Main Execution Stopped..")
}
