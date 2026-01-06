package main

import "fmt"

var ab = 20

func main() {
	fmt.Println(ab)

}

func init() {
	ab = 10
	//shadowing
	//ab := 10

	fmt.Println(ab)
}
