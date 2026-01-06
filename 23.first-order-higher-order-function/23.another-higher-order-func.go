package main

import "fmt"

func callHigherOrderFunc() func(x, y int) {
	return addVals
}

func addVals(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := callHigherOrderFunc()
	sum(100, 200)
}
