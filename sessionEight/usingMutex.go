/*

1. Creating Mutex
2. Restricting Developer to access some part of function concurrently

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var sampleMutex = &sync.Mutex{}
var waitGrp sync.WaitGroup

var sampleData = []int{5, 5, 5, 5, 5}

//var dataMutex = &sync.Mutex{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	waitGrp.Add(2)

	go AddDataToArray()
	go ReadDataFromArray()

	waitGrp.Wait()
}

func AddDataToArray() {
	//dataMutex.Lock()
	for i := 0; i < len(sampleData); i++ {
		sampleData[i] = sampleData[i] + 5
	}
	//dataMutex.Unlock()
	waitGrp.Done()
}

func ReadDataFromArray() {
	//dataMutex.Lock()
	for i := 0; i < len(sampleData); i++ {
		fmt.Println(sampleData[i])

	}
	//dataMutex.Unlock()
	waitGrp.Done()
}
