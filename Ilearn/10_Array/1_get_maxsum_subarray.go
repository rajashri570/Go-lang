//write a program to get the subarray with max sum

package main

import "fmt"

func main() {

	var arr = [7]int{-2, -3, 4, -1, -2, 1, 5}
	//var i int
	var maxsum = 0
	var startind int
	var endind int

	for i := 0; i < len(arr); i++ {

		for j := i; j < len(arr); j++ {
			var sum = 0
			fmt.Println(arr[i : j+1])
			for _, k := range arr[i : j+1] {
				sum = sum + k
			}
			if sum > maxsum {
				maxsum = sum
				startind = i
				endind = j + 1

			}

		}
		fmt.Println()

	}
	fmt.Println("Maxsum : ", maxsum)
	fmt.Println("Sub Array: ", arr[startind:endind])

}
