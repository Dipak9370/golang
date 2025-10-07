package main

import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	number := 8 // Change this value to test other numbers
	if isEven(number) {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}
}
