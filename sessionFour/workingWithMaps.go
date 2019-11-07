/*

1. Creating Data Structure Map
2. Different Ways to add map (Short Hand Notation or using var)
3. Getting and Setting Value in Map
4. Accessing non existing key in maps
5. Unordered Implementation in Maps (Ranging over keys): for key, value := range someMap {}
6. Reference type or Value type (Pass to function and Update)
Deleting Key values (delete(someMap, "Name"))

*/

package main

import "fmt"

func main() {
	var mapStructure = map[string]string{
		"name": "Mayank",
		"age":  "30",
	}

	UpdateMap(mapStructure)

	fmt.Println(mapStructure["name"])

	fmt.Println(mapStructure)
}

func UpdateMap(mapData map[string]string) {
	mapData["name"] = "Anshul"
}
