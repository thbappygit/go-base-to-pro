package main

import "fmt"

func main() {
	age := 20
	sex := "Male"

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	if age > 18 && sex == "Male" {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	grade := 85

	switch {
	case grade >= 80:
		fmt.Println("A+")
	case grade >= 70:
		fmt.Println("A")
	case grade >= 60:
		fmt.Println("B")
	default:
		fmt.Println("Fail")
	}

}
