package main

import (
	"WrodMap/wordFreqMap"
	"fmt"
)

func main() {
	str := "i am rajashri.I am woring in ambersoft. rajashri working since years.she loves programming. she try her best"
	fmt.Println(wordFreqMap.GetFreq(str))
}
