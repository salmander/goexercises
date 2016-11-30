package main

import "fmt"

func main() {
	// Function returns int
	fmt.Println("2 * 2 =", one(2))

	// Function returns int and string
	number, word := two(2, "Someword")
	fmt.Printf("number: %v, string: %v\n", number, word)

	// Naked return function
	fmt.Println("I will always return:", three())
}

func one(x int) int {
	fmt.Println("Function 1")
	return x * x
}

func two(number int, word string) (int, string) {
	fmt.Println("Function 2")
	return number, word
}

func three() (secret int) {
	fmt.Println("Function 3")
	return 3
}
