package main

import "fmt"

func sayMessage() (message string) {
	message = "Hello World"
	fmt.Println(message)

	showNewMessage := func() {
		message = "Hello Golang"
		fmt.Println("defer", message)
	}

	defer showNewMessage()

	message = "Hello Universe"
	fmt.Println("last", message)
	return

}
func expenses() int {
	ex := 100
	ex2 := 200
	newExpense := func() {
		ex = ex + ex2
	}
	defer newExpense()
	ex = 50

	return ex

}

func main() {

	sayMessage()
	fmt.Println(expenses())
	var data = [5]int{1, 2, 3, 4, 5}
	for _, v := range data {
		func() {
			defer fmt.Println(v)
		}()
	}

}
