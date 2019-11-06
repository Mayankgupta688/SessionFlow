/*

1. Adding Functions to the Struct
2. Updating the value of Objects in Structs
3. Creating Getters and Setters in Struct

*/

package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	Name string
	Age  int
}

func (emp Employee) GetEmployeeDetails() string {
	return emp.Name + " " + strconv.Itoa(emp.Age)
}

func main() {
	newEmp := Employee{"Mayank", 30}
	fmt.Println(newEmp.GetEmployeeDetails())
}
