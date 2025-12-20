package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)

}

func changeSlice(sliceNum []int) []int {

	sliceNum[0] = 10
	return append(sliceNum, 11)
}
