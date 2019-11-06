/*

1. Creating Private Variables
2. Creating Public Properties and Function
3. Package Level Private and Public Access
4. Creating Getters and Setters

*/

package publicprivateproperties

import "fmt"

type Employee struct {
	Name        string
	Age         int
	designation string
}

func AccessEmployeePublic() {
	accessEmployeePrivate()
}

func accessEmployeePrivate() {
	newEmployee := Employee{"Mayank", 10, "Developer"}
	fmt.Println(newEmployee.Name)
}
