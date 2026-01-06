package main

import "fmt"

// lets write a higher order function that takes a function as parameter
func processOperation(x int, y int, operation func(m, n int)) {
	operation(x, y)
}

func addCalc(p, q int) { /*who receive two int parameters*/
	r := p + q
	fmt.Println(r)
}

func main() {
	//addCalc(10,20)    /*who sent two int argument*/
	processOperation(40, 60, addCalc)
}

func init() {
	fmt.Println("Welcome to Higher Order Functions in GoLang")
	func(a int, b int, c int) {
		fmt.Println(a, b, c)
	}(1, 2, 3)
}
