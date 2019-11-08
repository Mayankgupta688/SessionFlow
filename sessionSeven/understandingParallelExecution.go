package main

import (
	"SessionFlow/data"
	"fmt"
)

func main() {
	ShowInitialData()
}

func ShowInitialData() {
	empData := data.GetEmployeeData()
	for i := 0; i < len(empData); i++ {
		fmt.Println(empData[i].Name)
	}
}
