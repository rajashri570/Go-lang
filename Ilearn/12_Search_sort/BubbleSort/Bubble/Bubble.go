package Bubble

type CompareFunc func(a, b rune) bool

func Bubble_sort(str []rune, compareFunc CompareFunc) []rune {
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if compareFunc(str[j], str[j+1]) {
				temp := str[j]
				str[j] = str[j+1]
				str[j+1] = temp
			}
		}
	}
	return str

}

/*

func main() {
	// var name = "rajashri"
	// str := []rune(name)
	var input_str string
	var ch int

	//annonymous function for ascending order
	asc := func(a, b rune) bool {
		return a > b
	}

	//annonymous function for Descending order
	des := func(a, b rune) bool {
		return a < b
	}
	for {
		fmt.Println("-------Menu---------")
		fmt.Println("1.Ascending Order")
		fmt.Println("2.Descending Oredr")
		fmt.Println("3.Quit")
		fmt.Println("Enter your choice :")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			fmt.Println("Enter the string : ")
			fmt.Scan(&input_str)
			str := []rune(input_str)
			arr := bubble_sort(str, asc)
			fmt.Println("sorted desc string:", string(arr))
		case 2:
			fmt.Println("Enter the string : ")
			fmt.Scan(&input_str)
			str := []rune(input_str)
			arr := bubble_sort(str, des)
			fmt.Println("sorted desc string :", string(arr))
		case 3:
			os.Exit(0)

		default:
			fmt.Println("invalid choice")

		}
	}

}
*/
