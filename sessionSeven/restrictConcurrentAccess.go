/*

Focuss on how to restrict the data to accessed concurrently by Multiple Goroutines
1. Create Global Data
2. Add multiple Gorouting to the Application
3. Add data to shared variable
4. Restrict Data to be used Concurrently by Using Channels
5. Restrict Access between channel input and output

*/

package main

import (
	"SessionFlow/data"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var internalData = data.GetEmployeeData()

func main() {

	ch := make(chan int, 1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	go AddEmployeeData(ch)
	go AddEmployeeData(ch)
	go AddEmployeeData(ch)

	time.Sleep(time.Second * 10)
}

func AddEmployeeData(ch chan int) {
	processId := strconv.Itoa(rand.Intn(100) + rand.Intn(400))
	ch <- 1
	for i := 0; i < len(internalData); i++ {
		fmt.Println(processId + " " + internalData[i].Name)
	}

	<-ch
}
