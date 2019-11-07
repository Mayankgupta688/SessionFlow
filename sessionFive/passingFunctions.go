package main

import "fmt"

type Employee interface {
	GetUserDetails()
}

type Manager struct {
	Name string
	Age  int
}

type Lead struct {
	Name     string
	TeamSize int
}

func (mg Manager) GetUserDetails() {
	fmt.Println("This is the Manager ", mg.Name)
}

func (ld Lead) GetUserDetails() {
	fmt.Println("This is the Lead ", ld.Name)
}

func main() {
	var emp Employee

	var leadData = Lead{"Mayank", 10}
	var managerData = Manager{"Mayank", 10}

	// Intyerface can have the assignment to both Manager and Lead Object...
	emp = managerData
	UpdateName(emp)

	emp = leadData
	UpdateName(emp)
}

func UpdateName(emp Employee) {
	emp.GetUserDetails()
}
