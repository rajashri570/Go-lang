// write a program to find the number of cnsonants in given string .
//input : rajashri
// vov - 3

package main

import (
	"fmt"
	"strings"
)

func main() {
	vovels := "aeiouAEIOU"

	str := "rutvik"

	vovel_cnt, conso_cnt := 0, 0

	for _, v := range str {
		if strings.ContainsRune(vovels, v) {
			vovel_cnt = vovel_cnt + 1
		} else {
			conso_cnt = conso_cnt + 1
		}
	}
	fmt.Println("No of vovels :", vovel_cnt, "\nNo of Consonant :", conso_cnt)

}
