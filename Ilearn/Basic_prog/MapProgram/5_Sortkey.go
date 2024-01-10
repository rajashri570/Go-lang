//write a programs to sort the keys in map
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//created empty map
	var mapvar = map[string]int{}

	//created slice to store the keys
	var keys []string

	//assigning values to map

	mapvar["rajashri"] = 26
	mapvar["Shrutika"] = 18
	mapvar["janu"] = 22
	mapvar["aparna"] = 15

	fmt.Println("----Before sorting map values--")
	for key, value := range mapvar {
		fmt.Println(key, value)
	}

	//ireating through for loop and appending key to slice

	for key := range mapvar {

		keys = append(keys, strings.ToLower(key))
	}
	sort.Strings(keys)
	fmt.Println("---After Sorting Map----")
	for _, key := range keys {
		fmt.Println(key, mapvar[key])
	}

	//fmt.Println(mapvar)
}

/*
----Before sorting map values--
janu 22
aparna 15
rajashri 26
Shrutika 18
---After Sorting Map----
aparna 15
janu 22
rajashri 26
shrutika 0
*/
