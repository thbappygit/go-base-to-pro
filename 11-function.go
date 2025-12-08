package main

import "fmt"

func getCalculationsOfTwoNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func printName(name string) {
	fmt.Println(name + "\n")
	fmt.Println(name + " is a good boy")
}

func main() {
	a := 15
	b := 20
	sum, mul := getCalculationsOfTwoNumbers(a, b)
	fmt.Println(sum)
	fmt.Println(mul)
	printName("Tanvir Hossen Bappy")
}
