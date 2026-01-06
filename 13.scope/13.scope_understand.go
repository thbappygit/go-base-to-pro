package main

import "fmt"

var (
	a = 20
	b = 30
)

func sumNumbers(x int, y int) {
	z := x + y
	fmt.Println(z)

}

func main() {
	p := 30
	q := 40

	sumNumbers(p, q)
	sumNumbers(a, q)
	sumNumbers(a, p)
}
