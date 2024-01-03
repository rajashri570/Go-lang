/*Problem Statement: Given a string s and a non-empty string p,
find all the starting indices of anagrams of p in s.

Example: plaintext Copy code Input: s = "cbaebabacd", p = "abc" Output: [0, 6] Explanation: The substring with starting index 0 is "cba," which is an anagram of "abc."
The substring with starting index 6 is "bac," which is an anagram of "abc
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "cbaebabacd"
	search := "abc"
	sub := []string{}

	for i := 0; i < len(str)-3; i++ {
		// created posiible substrings with len 3 and appended to sub slice
		sub = append(sub, str[i:i+3])

	}
	//fmt.Print(sub)
	//iterate through all slice sub

	for i, ele := range sub {
		//sorting ele in slice and then compared with search str
		sortedEle := sortString(ele)
		// print index if found anagram string
		if search == sortedEle {
			fmt.Print(i, " ")
		}
	}

}
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
