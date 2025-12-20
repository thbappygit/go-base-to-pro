package main

import "fmt"

func printVar(nums ...int) {
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func main() {
	printVar(5, 5, 6, 7, 8, 11, 15)
}
