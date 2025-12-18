package main

import "fmt"

var arrOfColegues = [5]string{"Bappy", "Emon", "Miraz", "Nasim", "Forhad"}
var justice = [2]bool{true, false}

func init() {
	fmt.Println("Welcome to array learning!")
}

func main() {
	fmt.Println(justice)

	var arrNumbers [5]int
	fmt.Println(arrOfColegues[1], "and", arrOfColegues[0], "are in same team")

	fmt.Println(arrNumbers[0], arrNumbers[4])

	length := len(arrNumbers)
	fmt.Println(length)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}
