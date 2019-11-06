/*

1. Composing Multiple Structs Together in one struct
2. Polymorphism in Golang Function
3. Concept of Object Composition vs Object Inheritance
4. Composing Struct having same function Name
5. Overriding Composed Level Function in the Parent Function

*/

package main

import "fmt"

type Employee struct {
	Name        string
	Age         int
	Designation string
}

func (empo Employee) getDetails() {
	fmt.Println(empo.Designation + " " + empo.Name)
}

type Manager struct {
	Employee
	TeamSize int
}

func main() {
	newManager := Manager{Employee{"Mayank", 12, "Developer"}, 19}
	newManager.Name = "Mayank"
	newManager.getDetails()
}
