package main

import "fmt"

func main() {

	//created nil map
	var map_1 map[int]int
	if map_1 == nil {
		// Checking if the map is nil or not

		//map_1[1] = 23

		fmt.Println("Map1 is nil")
	} else {

		fmt.Println("map1 is not nil")
	}

	map_1 = map[int]int{}
	if map_1 == nil {
		fmt.Println("map_1 is nil map")
	} else {
		fmt.Println("map1 after map[int]int{} not nil map..\nyou element to it")
	}
	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals

}
