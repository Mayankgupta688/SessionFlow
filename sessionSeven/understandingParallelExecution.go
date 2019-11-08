/*

1. Understanding Synchronous Execution
2. Execute Multiple Goroutines on the single core
3. Explain concept of Concurrency
4. Add Processors to the Application with "runtime.GOMAXPROCS"
5. See number of CPUs with "runtime.NumCPU()"
6. Run application in multiple cores..

*/

package main

import (
	"SessionFlow/data"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	go ShowInitialData()
	go ShowInitialDataStudent()
	go ShowInitialData()
	go ShowInitialDataStudent()
	go ShowInitialData()
	go ShowInitialDataStudent()
	time.Sleep(4 * time.Second)
}

func ShowInitialData() {
	empData := data.GetEmployeeData()
	for i := 0; i < len(empData); i++ {
		fmt.Println(empData[i].Name)
	}
}

func ShowInitialDataStudent() {
	studData := data.GetStudentData()
	for i := 0; i < len(studData); i++ {
		fmt.Println(studData[i].Name)
	}
}
