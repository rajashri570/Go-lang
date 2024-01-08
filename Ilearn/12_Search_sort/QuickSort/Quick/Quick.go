package Quick

func Quick_sort(arr []float64) []float64 {

	if len(arr) <= 1 {
		return arr
	}

	pivot_ind := len(arr) / 2
	pivot := arr[pivot_ind]

	var left_subarr, right_subarr []float64

	for i := 0; i < len(arr); i++ {
		if i == pivot_ind {
			continue
		}
		if arr[i] <= pivot {
			left_subarr = append(left_subarr, arr[i])
		} else {
			right_subarr = append(right_subarr, arr[i])
		}
	}

	left_subarr = Quick_sort(left_subarr)
	right_subarr = Quick_sort(right_subarr)

	// concatenate both arrays with pivot in the middle
	result := append(append(left_subarr, pivot), right_subarr...)

	return result
}

/*
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
	sorted := Quick_sort(float_arr)
	end := time.Now()
	elapse := end.Sub(start)
	fmt.Println("\nAfter sorting: ", sorted)

	fmt.Println("\nTine required for Quicksort :", elapse)

}
*/
