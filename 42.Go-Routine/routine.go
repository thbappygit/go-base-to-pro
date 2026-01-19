package main

import (
	"fmt"
	"time"
)

func stringFunc(text string) {
	for i := range 3 {
		fmt.Println(text, ":", i)
	}
}

func main() {
	stringFunc("Normal Function")

	go stringFunc("Go Routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
