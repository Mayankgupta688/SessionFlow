/*

1. Creating Interface
2. Adding Methods to Interface
3. Creating Structure with same Methods
4. Assigning Structure to Interface

*/

package main

import "fmt"

type EmployeeInterface interface {
	GetEmployeeDetails()
}

type Manager struct {
	Name string
	Age  int
}

func (mg Manager) GetEmployeeDetails() {
	fmt.Println("Employee Name is: " + mg.Name)
}

func main() {
	var sampleInterface EmployeeInterface
	var managerData = Manager{Name: "Mayank", Age: 20}
	sampleInterface = managerData
	sampleInterface.GetEmployeeDetails()
}
