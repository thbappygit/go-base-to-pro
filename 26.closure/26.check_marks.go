package main

import "fmt"

func checkMarks(num int) {
	marks := num
	if marks >= 40 {
		totalMarks := marks / 100
		fmt.Println("Status : Pass")
		fmt.Println("Total Marks :", totalMarks, "%")
	}
}

func main() {

	checkMarks(100)

	fmt.Println("Thank you for learning Go Programming")
}
