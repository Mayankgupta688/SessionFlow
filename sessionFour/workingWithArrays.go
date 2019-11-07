/*

1. Creating Simple Arrays with Golang
2. List of Values of different types... not possible
3. Passing Arrays to the function
4. Iteration in Array.. index, val := rangeArrayObject
5. Initial Values of Array properties..
6. Array without difining initial size in declaration...
7. Length is same as that used in initialization..

*/

package main

import "fmt"

func main() {

	// Defining Array
	arrayData := [8]int{1, 2, 3, 4}

	// Initial Value of the non initialized elements
	fmt.Println(arrayData[7])

	// Finding Length of Array
	fmt.Println(len(arrayData))

	//Iterating Arrays
	for _, data := range arrayData {
		fmt.Println(data)
	}

	// Initialize Array without specifying length
	randomArray := [...]int{1, 2, 3, 4}
	fmt.Println("Length of Random Array: ", len(randomArray))

	// Passing Array by Value.. (Default Behavior)
	UpdateArrayData(randomArray)
	fmt.Println("Value after updates: ", randomArray[1])

}

func UpdateArrayData(arr [4]int) {
	arr[1] = 10
	fmt.Println("Value before Update: ", arr[1])
}
