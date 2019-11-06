/*

1. Try accessing the private property in the Structure
2. Try Accessing Private Functions in the Packege

*/



package main

import (
	"SessionFlow/sessionThree/publicprivateproperties"
	"fmt"
)

func main() {
	newEmployee := new(publicprivateproperties.Employee)
	newEmployee.Name = "Mayank"
	newEmployee.designation = "Developer"
	publicprivateproperties.AccessEmployeePublic()
	fmt.Println(*newEmployee)
}
