package main

import (
	"strings"
)

func GetFreq(str string) map[string]int {

	wordmap := make(map[string]int)
	str2 := strings.ReplaceAll(str, ".", " ")
	// fmt.Println(str2)
	wordslice := strings.Fields(str2)
	for _, word := range wordslice {

		wordmap[word]++
	}

	return wordmap
}

func main() {

	str := "i am rajashri.I am woring in ambersoft. rajashri working since years.she loves programming. she try her best"

}
