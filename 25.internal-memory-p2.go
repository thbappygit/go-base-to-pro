package main

import "fmt"

const age = 25

var p = 30
var r = 10

func callN(c, d int) {
	e := c + d
	fmt.Println(e)
}

func main() {
	add := func(a, b int) {
		fmt.Println(a + b)
	}

	add(age, r)
	callN(age, p)
}

func init() {
	fmt.Println("Welcome to Go Programming")
}
