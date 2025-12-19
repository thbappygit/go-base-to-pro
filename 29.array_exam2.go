package main

import "fmt"

var arrData = [5]int{1, 2, 3, 4, 5}

func main() {

	arrData[0] = 2

	var command string
	var command2 string

	for _, value := range arrData {
		fmt.Println("I love you", value, "<")
	}

	for i := arrData[1]; i < arrData[0]; i++ {

		command, command2 = "I love you", "\n"

		fmt.Println(command, command2)

	}

	helloMyArrData := [10]string{"Hello", "My", "Array", "Data", "Lorem", "Ipsum", "Dolor", "Sit", "Amet", "Love"}

	fmt.Println(len(helloMyArrData))
	fmt.Println(cap(helloMyArrData))

}

func init() {
	fmt.Println("\"Welcome to Go Programming\"")
}
