/*

1. Create Start Function with Block Name
2. Create Stop Function with Block Name
3. Find Time Duration of Execution "time.Now()"
4. Using "Sub" function to substarct 2 times..
5. Create Simple Tracer Function
6. Complex Tracer Function

*/

package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	defer Stop(Start("Sample"))
	fmt.Println("Starting Function")
	fmt.Println("Stoping Function")
	fmt.Println(reflect.TypeOf(time.Now()))
}

func Start(blockName string) (string, time.Time) {
	var startTime = time.Now()
	fmt.Println("Started ", blockName+"  at: ", startTime)
	return blockName, startTime
}

func Stop(blockName string, startTime time.Time) {
	fmt.Println("Stoping Block: ", blockName, " at: ", time.Now().Sub(startTime))
}
