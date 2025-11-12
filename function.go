package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	fmt.Println("Sum of two numbers is: ", sum)
	return 1
}

func main() {
	add(10, 20)
}
