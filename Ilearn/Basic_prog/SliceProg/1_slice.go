/*write aprogram to create slice in go with diff way
- Understanding :
1. length specifies the number of element we want means in mycaseits 5 but
	but til i tried to adding two more element the cap was 7 means we can add upto 7 elements but the point
	i added 8th element the capacity got doubled initialy it was 7 after 8th elemnt became 14
2.  but the length = no elements in slice

*/
package main

import "fmt"

func main() {
	//first way

	var slice1 = make([]int, 5, 7)

	fmt.Println(slice1)

	fmt.Printf("Enter the %d element for slice:", len(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Scan(&slice1[i])

	}
	fmt.Println("Slice : ", slice1, "\nSlice Capacity: ", cap(slice1), "\nSlice LEngth:", len(slice1))
	slice1 = append(slice1, 23)
	fmt.Println("Slice : ", slice1, "\nSlice Capacity: ", cap(slice1), "\nSlice LEngth:", len(slice1))
	slice1 = append(slice1, 10)
	fmt.Println("Slice : ", slice1, "\nSlice Capacity: ", cap(slice1), "\nSlice LEngth:", len(slice1))
	slice1 = append(slice1, 90)
	fmt.Println("Slice : ", slice1, "\nSlice Capacity: ", cap(slice1), "\nSlice LEngth:", len(slice1))
}

/*
rdolzake@rdolzake-HP-245-G5-Notebook-PC:~/GOLANG/Go-lang/Ilearn/Basic_prog$ go run 1_slice.go
Enter the 5 element for slice: 12 23 44 11 22

Slice :  [12 23 44 11 22]
Slice Capacity:  7
Slice LEngth: 5
Slice :  [12 23 44 11 22 23]
Slice Capacity:  7
Slice LEngth: 6
rdolzake@rdolzake-HP-245-G5-Notebook-PC:~/GOLANG/Go-lang/Ilearn/Basic_prog$ go run 1_slice.go
Enter the 5 element for slice11 12 13 14 15
Slice :  [11 12 13 14 15]
Slice Capacity:  7
Slice LEngth: 5
Slice :  [11 12 13 14 15 23]
Slice Capacity:  7
Slice LEngth: 6
Slice :  [11 12 13 14 15 23 10]
Slice Capacity:  7
Slice LEngth: 7
Slice :  [11 12 13 14 15 23 10 90]
Slice Capacity:  14
Slice LEngth: 8
*/
