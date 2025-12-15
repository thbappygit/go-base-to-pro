package main

import "fmt"

func higherFunc(x, y int, amarvai func(a, b int)) {
	amarvai(x, y)
}

func emon(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	higherFunc(10, 20, emon)
}
