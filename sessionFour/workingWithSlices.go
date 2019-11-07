/*

1. Understanding Difference between length and capacity
2. Slices are memory locations
3. Slices are passed by reference
4. Slice can be a sub part of an array
5. Output for "fmt.Println" for slice is not starting memory reference of the Slice element
6. Append to Slices

*/

package main

import "fmt"

func main() {

	// Undersstanding Capacity and Length Difference...
	sliceData := []int{1, 2, 3}
	fmt.Println(cap(sliceData))
	fmt.Println(len(sliceData))

	sliceData = append(sliceData, 4)
	fmt.Println(cap(sliceData))
	fmt.Println(len(sliceData))

	sliceData = append(sliceData, 5)
	sliceData = append(sliceData, 6)
	fmt.Println(len(sliceData))

	sliceData = append(sliceData, 7)
	sliceData = append(sliceData, 8)
	fmt.Println(cap(sliceData))
	fmt.Println(len(sliceData))

	UpdateSlice(sliceData)
	fmt.Println(sliceData[0])
	fmt.Println(sliceData)
}

func UpdateSlice(sliceData []int) {
	sliceData[0] = 20
}
