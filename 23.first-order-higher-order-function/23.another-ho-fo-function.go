package main

import "fmt"

func myOperateFunction(a, b int, op func(x, y int) int) {
	res := op(a, b)
	fmt.Println(res)
}

func addFunc(x, y int) int {
	return x + y
}

func subFunc(x, y int) int {
	return x - y
}

func mulFunc(x, y int) int {
	return x * y
}
func main() {

	myOperateFunction(10, 30, addFunc)
	myOperateFunction(89, 30, subFunc)
	myOperateFunction(90, 3, mulFunc)
}
