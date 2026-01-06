package main

import "fmt"

func main() {
	var a int8 = 127
	var b int8 = -128

	var c uint8 = 255 //uint means unsigned (0 or positive integers)
	var d uint16 = 1000

	var j float32 = 10.2242424
	var k float64 = 29.56 //If your computer is a 32 bit computer then it will take 2 cell of memory 1 for 32 and other for 32 to initiate the variable into code segment

	var flag bool = false

	love := '❤' //rune datatype is used to store a single character. It is an alias for int32

	fmt.Println(a, b, c, d, j, k, flag, love)

	fmt.Printf("%c\n", love) //to print rune data
	fmt.Printf("%d\n", a)    //to print int data
	fmt.Printf("%.3f\n", j)  //to print float32 data .3 is print after 3 numbers after point
	fmt.Printf("%v\n", flag)
	fmt.Printf("%T\n", flag) //to print type
}
