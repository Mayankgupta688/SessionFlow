package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name        string
	Designation string
}

func main() {
	newEmployee := Employee{"Mayank", "Desveloper"}
	ch := make(chan Employee, 1)
	go GetEmployeeData(ch)
	ch <- newEmployee
	time.Sleep(5 * time.Second)
}

func GetEmployeeData(ch chan Employee) {
	empData := <-ch
	fmt.Println(empData.Name)
}
