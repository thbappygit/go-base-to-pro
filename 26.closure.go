package main

import "fmt"

const abc = 10

var amt = 50

func outer() func() {
	money := 100
	myAge := 20
	fmt.Println("Age is :", myAge)

	show := func() {
		money = money + abc + amt
		fmt.Println("Money is :", money)
	}
	return show
}

func call() {
	inner := outer()
	inner()
	inner()

	inner2 := outer()
	inner2()
	inner2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Welcome to closure learning!")
}
