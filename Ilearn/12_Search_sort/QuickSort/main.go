package main

import (
	"QuickSort/Quick"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min_range := 1.5
	max_range := 10.0

	var float_arr []float64

	for i := 0; i < 5; i++ {

		rand.Seed(time.Now().UnixNano())
		randnum := min_range + rand.Float64()*(max_range-min_range)
		float_arr = append(float_arr, randnum)

	}
	fmt.Println("Before sorting: ", float_arr)
	start := time.Now()
	sorted := Quick.Quick_sort(float_arr)
	end := time.Now()
	elapse := end.Sub(start)
	fmt.Println("\nAfter sorting: ", sorted)

	fmt.Println("\nTine required for Quicksort :", elapse)

}
