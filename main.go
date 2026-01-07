package main

import "fmt"

var arrOfColleagues = [5]string{"Bappy", "Emon", "Miraz", "Nasim", "Forhad"}
var justice = [2]bool{true, false}

func init() {
	fmt.Println("Welcome to array learning!")
}

func main() {
	fmt.Println(justice)

	var arrNumbers [5]int
	fmt.Println(arrOfColleagues[1], "and", arrOfColleagues[0], "are in same team")

	fmt.Println(arrNumbers[0], arrNumbers[4])

	length := len(arrNumbers)
	fmt.Println(length)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{7, 8, 9, 10, 11}
	fmt.Println("dcl2:", b)

	for index, value := range b {
		fmt.Println(index, value)
	}

}
