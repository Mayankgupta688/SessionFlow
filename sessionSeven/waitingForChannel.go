/*

1. Create a Main Function
2. Add a Goroutine
3. Pass channel
4. Add Data in Channel in Gorouting
5. Consume data from the channel
6. Create ConsumerFunction and ProducerFunction for Channel..

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main Execution Block Starts")

	ch := make(chan int)

	go RoutingProducer(ch)

	go RoutingConsumer(ch)

	time.Sleep(5 * time.Second)

	fmt.Println("Main Execution Block Stops")
}

func RoutingProducer(ch chan int) {
	ch <- 10
}

func RoutingConsumer(ch chan int) {
	fmt.Println(<-ch)
}
