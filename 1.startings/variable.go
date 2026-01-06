package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100
	const name = "Tanvir"
	fmt.Println(name)
	fmt.Printf("The value of a is %v and its type is %v\n\n", a, reflect.TypeOf(a))
}
