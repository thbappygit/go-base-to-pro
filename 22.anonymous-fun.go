package main

import "fmt"

func main() {
	//Immediately Invoked Function Expression OR IIFE function or Anonymous function
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20)

	fmt.Println("Hello")
}

func init() {
	fmt.Println("Init")
}
