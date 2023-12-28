package main

import "fmt"

func greet() {
	fmt.Println("Good morning!")
}
func main() {
	name := "rajashri"
	defer fmt.Println(name)
	defer fmt.Println(name + " dolzake")
	greet()

}
