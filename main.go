package main

import (
	"fmt"
	"time"
)

func addCalc(a, b int) {
	fmt.Println(a + b)
}

func printHello(num int) {
	fmt.Println("Hello Bappy", num)
	addCalc(10, 20)
}

func main() {
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

	time.Sleep(4 * time.Second)
}
