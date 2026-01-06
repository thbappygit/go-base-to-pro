package main

import "fmt"

func main() {
	examMark := 100
	addressOfExamMark := &examMark // this is the variable to take the address/pointer of examMark Variable // & is pointer operator

	valueOfTheAddress := *addressOfExamMark

	fmt.Println("Address of examMark : ", addressOfExamMark, "")
	fmt.Println("value of examMark : ", valueOfTheAddress, "")
}
