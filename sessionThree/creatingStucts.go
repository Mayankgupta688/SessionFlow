/*

1. Creating Simple Stuct in Golang Ëmployee
2. Add Properties to the Struct
3. Creating Object for Struct with Employee{}
6. Not passing all the values while creating object with Employee{}
4. Creating Object with "new" keyword
5. Default values of Fields in Struct when using "new"
7. Impact of returning Address in "new" keyword
8. Zeroed Value of different data types

*/

package main

import "fmt"

type Employee struct {
	Name        string
	Designation string
	Age         int
}

func main() {

	// Method One...

	newEmp := Employee{"Mayank", "Developer", 30}

	UpdateName(newEmp)

	fmt.Println(newEmp.Name)

	otherEmployee := new(Employee)

	UpdateReferenceName(otherEmployee)

	fmt.Println(otherEmployee.Name)
}

func UpdateName(emp Employee) {
	emp.Name = "Meha"
}

func UpdateReferenceName(emp *Employee) {
	emp.Name = "Meha"
}
