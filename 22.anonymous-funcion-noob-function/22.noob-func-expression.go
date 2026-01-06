package main

import "fmt"

func init() {
	fmt.Println("Init")
}

var myCounterFunc = func(maxNumb int) {
	for i := 0; i <= maxNumb; i++ {
		fmt.Println(i)
	}
}

func main() {
	noob := func(b, c int) {
		fmt.Println(b + c)
	}

	noob(10, 20)
	myCounterFunc(100)
}
