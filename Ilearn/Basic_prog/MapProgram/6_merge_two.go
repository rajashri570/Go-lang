//write a program to merger second map into first map

package main

import "fmt"

func main() {

	map1 := map[int]string{1: "raj", 2: "smrudhi", 3: "jaya"}
	map2 := map[int]string{4: "rachana", 2: "smrudhi", 6: "jaya", 7: "mayuri"}

	i := 0
	for range map1 {
		i += 1
	}
	for key, val := range map2 {
		map1[key] = val
	}

	fmt.Println(map1)

}
