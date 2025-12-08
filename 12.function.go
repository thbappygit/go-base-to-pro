package main

import "fmt"

func welcomeMsg() {
	fmt.Println("Welcome to Go Programming")
}
func enterName() string {
	var name string
	fmt.Println("Enter your name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	return name
}

func getTwoNumbers() (int, int) {
	fmt.Println("Enter 1st number:")
	var num1 int
	_, err := fmt.Scanln(&num1)
	if err != nil {
		return 0, 0
	}
	fmt.Println("Enter 2nd number:")
	var num2 int
	_, err = fmt.Scanln(&num2)
	if err != nil {
		return 0, 0
	}

	return num1, num2
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func sayGoodbye() {
	fmt.Println("Thank you for using this program.")
}

func main() {
	welcomeMsg()
	name := enterName()
	num1, num2 := getTwoNumbers()
	sum := addNumbers(num1, num2)
	fmt.Printf("Hello, %s! The sum of %d and %d is %d.\n", name, num1, num2, sum)
	sayGoodbye()

}
